package comm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func PrintUrl(url string) error {
	body, err := getBytes(url)
	if err != nil {
		return err
	}

	err = printJson(body)
	if err != nil {
		return err
	}
	return nil
}

var httpClient = &http.Client{Timeout: httpClientTimeout}

func getBytes(url string) ([]byte, error) {
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		bodyText := string(body)
		if bodyText == "" {
			bodyText = "<empty>"
		}
		return nil, fmt.Errorf("Error %d (%s) al cridar a %s", response.StatusCode, bodyText, url)
	}
	return body, nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func printJson(body []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "    "); err != nil {
		return err
	}
	fmt.Println(prettyJSON.String())
	return nil
}

func OpenUrl(url string) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return errors.New("Unsupported platform")
	}
}
