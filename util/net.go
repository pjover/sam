package util

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

var httpClient = &http.Client{Timeout: httpClientTimeout}

const contentType = "application/json; charset=UTF-8"

func Get(url string) ([]byte, error) {
	response, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if isKo(response) {
		return processError(url, body, response)
	}
	return body, nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func isKo(response *http.Response) bool {
	return response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusBadRequest
}

func processError(url string, body []byte, response *http.Response) ([]byte, error) {
	bodyText := string(body)
	if bodyText == "" {
		bodyText = "<empty>"
	}
	return nil, fmt.Errorf("Error %d (%s) al cridar a %s", response.StatusCode, bodyText, url)
}

func PrintGet(url string) error {
	body, err := Get(url)
	if err != nil {
		return err
	}

	err = printJson(body)
	if err != nil {
		return err
	}
	return nil
}

func printJson(body []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, body, "", "    "); err != nil {
		return err
	}
	fmt.Println(prettyJSON.String())
	return nil
}

func GetType(url string, target interface{}) error {
	response, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer closeBody(response.Body)

	return json.NewDecoder(response.Body).Decode(target)
}

func Post(url string, data []byte) ([]byte, error) {
	response, err := httpClient.Post(url, contentType, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer closeBody(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if isKo(response) {
		return processError(url, body, response)
	}
	return body, nil
}

func PrintPost(url string, data []byte) error {
	body, err := Post(url, data)
	if err != nil {
		return err
	}

	err = printJson(body)
	if err != nil {
		return err
	}
	return nil
}

func OpenOnBrowser(url string) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return errors.New("unsupported platform")
	}
}
