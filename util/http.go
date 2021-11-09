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
	"os"
	"os/exec"
	"runtime"
)

const contentType = "application/json; charset=UTF-8"

func NewHttpClient() *http.Client {
	return &http.Client{Timeout: httpClientTimeout}
}

type HttpGetManager interface {
	Bytes(url string) ([]byte, error)
	PrettyJson(url string) (string, error)
	Type(url string, target interface{}) error
}

type SamHttpGetManager struct {
	httpClient *http.Client
}

func NewHttpGetManager() HttpGetManager {
	return SamHttpGetManager{
		httpClient: NewHttpClient(),
	}
}

func (s SamHttpGetManager) Bytes(url string) ([]byte, error) {
	response, err := s.httpClient.Get(url)
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

func (s SamHttpGetManager) PrettyJson(url string) (string, error) {
	body, err := s.Bytes(url)
	if err != nil {
		return "", err
	}
	return ToPrettyJson(body)
}

func ToPrettyJson(body []byte) (string, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, body, "", "    ")
	if err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func (s SamHttpGetManager) Type(url string, target interface{}) error {
	response, err := s.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer closeBody(response.Body)
	return json.NewDecoder(response.Body).Decode(target)
}

type HttpPostManager interface {
	Bytes(url string, data []byte) ([]byte, error)
	PrettyJson(url string, data []byte) (string, error)
	File(remoteUrl string, filePath string) (string, error)
}

type SamHttpPostManager struct {
	httpClient *http.Client
}

func NewHttpPostManager() HttpPostManager {
	return SamHttpPostManager{
		httpClient: NewHttpClient(),
	}
}

func (s SamHttpPostManager) Bytes(url string, data []byte) ([]byte, error) {
	response, err := s.httpClient.Post(url, contentType, bytes.NewBuffer(data))
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

func (s SamHttpPostManager) PrettyJson(url string, data []byte) (string, error) {
	body, err := s.Bytes(url, data)
	if err != nil {
		return "", err
	}
	return ToPrettyJson(body)
}

func (s SamHttpPostManager) File(remoteUrl string, filePath string) (string, error) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	response, err := s.httpClient.Post(remoteUrl, contentType, nil)
	if err != nil {
		return "", err
	}
	defer closeBody(response.Body)

	_, err = io.Copy(file, response.Body)
	defer closeFile(file)
	if err != nil {
		return "", err
	}

	return fmt.Sprint("Creat el fitxer ", filePath), nil
}

func closeFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Fatal(err)
	}
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
