package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

// TODO remove after migration

// Untar decompresses the source file to target file
func Untar(ctx context.Context, source, target string) error {

	if err := os.MkdirAll(target, 0600); err != nil {
		return err
	}

	tarExecutable := "tar"
	if runtime.GOOS == "windows" {
		tarExecutable = "C:\\opscode\\chef\\bin\\tar.exe"
	}
	cmd := exec.CommandContext(ctx, tarExecutable, "-xzf", source, "-C", target)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// CheckStandardStatus return error if status is not OK
func CheckStandardStatus(status int, response []byte) error {

	if status < 300 {
		return nil
	}

	// default, raw, not parsing
	message := string(response[:])

	var responseContent map[string]interface{}
	err := json.Unmarshal(response, &responseContent)
	if err == nil {
		if responseContent["errors"] != nil {
			message = ""
			for key, value := range responseContent["errors"].(map[string]interface{}) {
				subMessages := make([]string, len(value.([]interface{})))
				for i, v := range value.([]interface{}) {
					subMessages[i] = fmt.Sprint(v)
				}
				composedMsg := strings.Join(subMessages, ",")
				message = fmt.Sprintf("%s#%s:%s", message, key, composedMsg)
			}
		} else if responseContent["error"] != nil {
			message = responseContent["error"].(string)
		}
	}

	return fmt.Errorf("HTTP request failed: (%d) [%s]", status, message)
}

// FileExists checks file existence
func FileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
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

// RemoveDuplicates returns the slice removing duplicates if exist
func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}

// Contains evaluates whether s contains x.
func Contains(s []string, x string) bool {
	for _, n := range s {
		if x == n {
			return true
		}
	}
	return false
}
