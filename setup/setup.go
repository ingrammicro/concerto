// +build !solaris

package setup

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"path"
	"strings"
	"syscall"

	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/utils"

	"github.com/asaskevich/govalidator"
	"github.com/codegangsta/cli"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/net/html"
	"golang.org/x/net/publicsuffix"
)

// WebClient stores Web Client data
type WebClient struct {
	client *http.Client
	url    *url.URL
	csrf   string
	cookie *http.Cookie
}

// NewWebClient creates a new Web Client
func NewWebClient(endpoint string) (*WebClient, error) {
	transport := &http.Transport{}
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Transport: transport, Jar: jar}

	endpointURL, err := url.ParseRequestURI(endpoint)
	if err != nil {
		return nil, err
	}

	return &WebClient{client, endpointURL, "", nil}, nil
}

func (w *WebClient) obtainCsrf(b io.Reader) error {
	var errorMessage error
	z := html.NewTokenizer(b)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return errorMessage
		case tt == html.SelfClosingTagToken:
			t := z.Token()
			isMeta := t.Data == "meta"
			if isMeta && len(t.Attr) > 0 {
				if (t.Attr[1].Key == "name") && (t.Attr[1].Val == "csrf-token") {
					w.csrf = t.Attr[0].Val
					log.Debugf("Csrf Token: %s", w.csrf)
				} else if (t.Attr[0].Key == "name") && (t.Attr[0].Val == "csrf-token") {
					w.csrf = t.Attr[1].Val
					log.Debugf("Csrf Token: %s", w.csrf)
				}
			}
		case tt == html.StartTagToken:
			t := z.Token()
			if (t.Data == "div") && len(t.Attr) > 0 && (t.Attr[0].Key == "id") && (t.Attr[0].Val == "flash_alert") {
				z.Next()
				errorMessage = errors.New(z.Token().String())
			}
		}
	}

}

func (w *WebClient) login(email string, password string) error {
	w.url.Path = "/accounts/login"

	response, err := w.client.Get(w.url.String())
	if err != nil {
		log.Fatalf("%#v", err)
	}
	defer response.Body.Close()

	err = w.obtainCsrf(response.Body)
	if err != nil {
		log.Fatalf("%#v", err)
	}

	if w.csrf == "" {
		log.Debugf("Can not log into %s as %s", w.url.String(), email)
		return fmt.Errorf("Can not log into %s as %s", w.url.String(), email)
	}

	w.cookie = response.Cookies()[0]

	account := url.Values{}
	account.Set("authenticity_token", w.csrf)
	account.Set("account[email]", email)
	account.Set("account[password]", password)

	response, err = w.client.PostForm(w.url.String(), account)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = w.obtainCsrf(response.Body)
	if err == nil {
		log.Debugf("Logged in %s as %s", w.url.String(), email)
	}

	return err
}

func (w *WebClient) generateAPIKeys() error {
	w.url.Path = "/settings/api_key"

	emptyValue := []byte("{}")
	request, err := http.NewRequest("POST", w.url.String(), bytes.NewBuffer(emptyValue))
	if err != nil {
		return err
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-CSRF-TOKEN", w.csrf)

	response, err := w.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode >= 300 {
		return fmt.Errorf(fmt.Sprintf("We couldn't check for the existence of api keys at your account. Please try by loging to %s and generating manually through settings > accounts", w.url.String()))
	}
	return nil
}

func (w *WebClient) getAPIKeys() error {
	w.url.Path = "/settings/api_key.zip"

	response, err := w.client.Get(w.url.String())
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < 300 && response.Header.Get("Content-Type") == "application/zip" {

		config, err := utils.GetConcertoConfig()
		if err != nil {
			log.Debug(err.Error)
			return err
		}

		concertoFolderSSL := path.Join(config.ConfLocation, "ssl")
		if !config.IsHost {
			os.MkdirAll(path.Join(concertoFolderSSL, "private"), 0755)
			file, err := ioutil.TempFile(os.TempDir(), "api-key.zip")
			if err != nil {
				return err
			}
			defer file.Close()
			io.Copy(file, response.Body)

			err = utils.Unzip(file.Name(), concertoFolderSSL)
			defer os.Remove(file.Name())
			if err != nil {
				return err
			}
			log.Debugf("Unziped Api Keys in %s.\n", concertoFolderSSL)
			return nil

		}
		return errors.New("You are trying to overwrite server configuration. Please contact your administrator")
	}
	return fmt.Errorf("We are not able to download your API keys. Please try by loging to %s/settings/api_key.zip in your web navigator ", w.url.String())
}

func cmdSetupAPIKeys(c *cli.Context) error {
	var emailUnClean string
	var passwordUnClean []byte

	reader := bufio.NewReader(os.Stdin)
	config, err := utils.GetConcertoConfig()
	if err != nil {
		log.Fatalf("Error getting current configuration: %s", err.Error)
	}

	loginURL := config.ConcertoURL
	if err != nil {
		log.Fatalf("Error getting Concerto URL: %s", err.Error)
	}

	fmt.Printf("Using Concerto endpoint %s \n", loginURL)
	if c.IsSet("email") {
		emailUnClean = c.String("email")
	} else {
		fmt.Printf("Email: ")
		emailUnClean, _ = reader.ReadString('\n')
	}

	if c.IsSet("password") {
		passwordUnClean = []byte(c.String("password"))
	} else {
		fmt.Printf("Password: ")
		passwordUnClean, _ = terminal.ReadPassword(int(syscall.Stdin))
	}

	email := strings.TrimSpace(string(emailUnClean))
	password := strings.TrimSpace(string(passwordUnClean))
	fmt.Printf("\n")

	if govalidator.IsEmail(email) {
		client, err := NewWebClient(loginURL)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Logging into Concerto ...")
		err = client.login(email, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" OK\n")

		fmt.Printf("Checking/Generating API keys ...")
		err = client.generateAPIKeys()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" OK\n")

		fmt.Printf("Downloading API keys ...")
		err = client.getAPIKeys()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(" OK\n")

	} else {
		log.Fatalf("Email address %s is not a valid email", email)
	}
	return nil
}

// SubCommands return API Keys subcommands
func SubCommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "api_keys",
			Usage:  "Install Concerto Api Keys",
			Action: cmdSetupAPIKeys,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "email",
					Usage: "Email used to log into concerto",
				},
				cli.StringFlag{
					Name:  "password",
					Usage: "Password used to log into concerto",
				},
			},
		},
	}
}
