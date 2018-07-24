package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/codegangsta/cli"

	log "github.com/Sirupsen/logrus"
)

// TODO remove after migration

// CheckError checks if error is received
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Unzip uncompress the received file into the target path
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

// ScrapeErrorMessage returns scrapped response
func ScrapeErrorMessage(message string, regExpression string) string {

	re, err := regexp.Compile(regExpression)
	scrapped := re.FindStringSubmatch(message)

	if scrapped == nil || err != nil || len(scrapped) < 2 {
		// couldn't scrape, return generic error
		message = "Error executing operation"
	} else {
		// return scrapped response
		message = scrapped[1]
	}

	return message
}

// CheckReturnCode validates error code and message if needed
func CheckReturnCode(res int, mesg []byte) {
	if res >= 300 {

		message := string(mesg[:])
		log.Debugf("IMCO API response: %s", message)

		f := func(c rune) bool {
			return c == ',' || c == ':' || c == '{' || c == '}' || c == '"' || c == ']' || c == '['
		}

		// check if response is a web page.
		if strings.Contains(message, "<html>") {
			scrapResponse := "<title>(.*?)</title>"
			message = ScrapeErrorMessage(message, scrapResponse)
		} else if strings.Contains(message, "{\"errors\":{") {
			scrapResponse := "{\"errors\":(.*?)}"

			message = ScrapeErrorMessage(message, scrapResponse)
			result := strings.Split(message, ",")
			if result != nil && len(result) >= 1 {
				message = result[0]
			}
			// Separate into fields with func.
			fields := strings.FieldsFunc(message, f)
			message = strings.Join(fields[:], " ")

		} else if strings.Contains(message, "{\"error\":") {
			scrapResponse := "{\"error\":\"(.*?)\"}"
			message = ScrapeErrorMessage(message, scrapResponse)
		}
		// if it's not a web page or json-formatted message, return the raw message
		log.Fatal(fmt.Sprintf("HTTP request failed: (%d) [%s]", res, message))
	}
}

// CheckStandardStatus return error if status is not OK
func CheckStandardStatus(status int, mesg []byte) error {

	if status < 300 {
		return nil
	}

	message := string(mesg[:])

	f := func(c rune) bool {
		return c == ',' || c == ':' || c == '{' || c == '}' || c == '"' || c == ']' || c == '['
	}

	if strings.Contains(message, "{\"errors\":{") {
		scrapResponse := "{\"errors\":(.*?)}"

		message = ScrapeErrorMessage(message, scrapResponse)
		result := strings.Split(message, ",")
		if result != nil && len(result) >= 1 {
			message = result[0]
		}
		// Separate into fields with func.
		fields := strings.FieldsFunc(message, f)
		message = strings.Join(fields[:], " ")

	} else if strings.Contains(message, "{\"error\":") {
		scrapResponse := "{\"error\":\"(.*?)\"}"
		message = ScrapeErrorMessage(message, scrapResponse)
	}

	// if it's not a web page or json-formatted message, return the raw message
	return fmt.Errorf("HTTP request failed: (%d) [%s]", status, message)
}

// FileExists checks file existence
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// CheckRequiredFlags checks for required flags, and show usage if requirements not met
func CheckRequiredFlags(c *cli.Context, flags []string) {
	missing := ""
	for _, flag := range flags {
		if !c.IsSet(flag) {
			missing = fmt.Sprintf("%s\t--%s\n", missing, flag)
		}
	}

	if missing != "" {
		fmt.Printf("Incorrect usage. Please use parameters:\n%s\n", missing)
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(2)
	}
}

// RandomString generates a random string from lowercase letters and numbers
func RandomString(strlen int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, strlen)
	for i := range result {
		result[i] = chars[r.Intn(len(chars))]
	}
	return string(result)
}
