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

const contentType = "application/json; charset=UTF-8"

func getHttpClient() *http.Client {
	return &http.Client{Timeout: httpClientTimeout}
}

type HttpGetManager interface {
	GetBytes(url string) ([]byte, error)
	GetPrettyJson(url string) (string, error)
	GetType(url string, target interface{}) error
	GetPrint(url string) (string, error)
}

type SamHttpGetManager struct {
	httpClient *http.Client
}

func NewHttpGetManager() HttpGetManager {
	return SamHttpGetManager{
		getHttpClient(),
	}
}

func (s SamHttpGetManager) GetBytes(url string) ([]byte, error) {
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

func (s SamHttpGetManager) GetPrettyJson(url string) (string, error) {
	body, err := s.GetBytes(url)
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

func (s SamHttpGetManager) GetType(url string, target interface{}) error {
	response, err := s.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer closeBody(response.Body)
	return json.NewDecoder(response.Body).Decode(target)
}

func (s SamHttpGetManager) GetPrint(url string) (string, error) {
	_json, err := s.GetPrettyJson(url)
	if err != nil {
		return "", err
	}
	fmt.Println(_json)
	return _json, nil
}

type HttpPostManager interface {
	PostBytes(url string, data []byte) ([]byte, error)
	PostPrettyJson(url string, data []byte) (string, error)
	PostPrint(url string, data []byte) (string, error)
}

type SamHttpPostManager struct {
	httpClient *http.Client
}

func NewHttpPostManager() HttpPostManager {
	return SamHttpPostManager{
		getHttpClient(),
	}
}

func (s SamHttpPostManager) PostBytes(url string, data []byte) ([]byte, error) {
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

func (s SamHttpPostManager) PostPrettyJson(url string, data []byte) (string, error) {
	body, err := s.PostBytes(url, data)
	if err != nil {
		return "", err
	}
	return ToPrettyJson(body)
}

func (s SamHttpPostManager) PostPrint(url string, data []byte) (string, error) {
	_json, err := s.PostPrettyJson(url, data)
	if err != nil {
		return "", err
	}
	fmt.Println(_json)
	return _json, nil
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
