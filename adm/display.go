package adm

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func DisplayCustomer(customerCode int) error {
	url := fmt.Sprintf("http://localhost:8080/customers/%d", customerCode)
	return display(url)
}

func display(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer closeBody(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyText := string(body)
	if bodyText == "" {
		bodyText = "<empty>"
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error %d (%s) al cridar a %s", resp.StatusCode, bodyText, url)
	}
	fmt.Println(bodyText)
	return nil
}

func closeBody(body io.ReadCloser) {
	err := body.Close()
	if err != nil {
		log.Fatal(err)
	}
}
