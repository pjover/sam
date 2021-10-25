package comm

import (
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

func PrintUrl(url string) error {
	body, err := getBody(url)
	if err != nil {
		return err
	}
	printYaml(body)
	return nil
}

func getBody(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer closeBody(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		bodyText := string(body)
		if bodyText == "" {
			bodyText = "<empty>"
		}
		return nil, fmt.Errorf("Error %d (%s) al cridar a %s", resp.StatusCode, bodyText, url)
	}
	return body, nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func printYaml(json []byte) {
	y, err := yaml.JSONToYAML(json)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
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
