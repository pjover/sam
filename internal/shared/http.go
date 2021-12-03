package shared

import (
	"archive/zip"
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
	"path"
	"runtime"
	"strings"
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
	FileWithDefaultName(remoteUrl string, directory string) (string, error)
	File(remoteUrl string, directory string, filename string) (string, error)
	Zip(remoteUrl string, directory string) (string, error)
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

func (s SamHttpPostManager) FileWithDefaultName(remoteUrl string, directory string) (string, error) {
	response, err := s.httpClient.Post(remoteUrl, contentType, nil)
	if err != nil {
		return "", err
	}
	filename := extractDefaultName(response.Header.Get("Content-Disposition"))
	defer closeBody(response.Body)
	return writeFile(response.Body, directory, filename)
}

func extractDefaultName(contentDisposition string) string {
	split := strings.Split(contentDisposition, ";")
	filename := strings.Split(split[2], "\"")
	return filename[1]
}

func (s SamHttpPostManager) File(remoteUrl string, directory string, filename string) (string, error) {
	response, err := s.httpClient.Post(remoteUrl, contentType, nil)
	if err != nil {
		return "", err
	}
	defer closeBody(response.Body)
	return writeFile(response.Body, directory, filename)
}

func (s SamHttpPostManager) Zip(remoteUrl string, directory string) (string, error) {

	req, err := http.NewRequest("POST", remoteUrl, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/zip")
	response, err := s.httpClient.Do(req)
	defer closeBody(response.Body)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return "", err
	}

	// Read all the files from zip archive
	var sb strings.Builder
	for _, zipFile := range zipReader.File {
		fmt.Println("Reading file:", zipFile.Name)
		byteData, err := readZipFile(zipFile)
		if err != nil {
			sb.WriteString(fmt.Sprintf(" ❌ %s\n", zipFile.Name))
			log.Println(err)
			continue
		}
		_, err = writeFile(bytes.NewReader(byteData), directory, zipFile.Name)
		if err != nil {
			return "", err
		}
		sb.WriteString(fmt.Sprintf(" ✔️ %s\n", zipFile.Name))
	}

	return sb.String(), nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer func(f io.ReadCloser) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	return ioutil.ReadAll(f)
}

func writeFile(reader io.Reader, directory string, filename string) (string, error) {
	filePath := path.Join(directory, filename)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(file, reader)
	defer closeFile(file)
	if err != nil {
		return "", err
	}

	return filePath, nil
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
