package adm

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ghodss/yaml"
)

func DisplayCustomer(customerCode int) error {
	url := fmt.Sprintf("http://localhost:8080/customers/%d", customerCode)
	return display(url)
}

func DisplayProduct(productCode string) error {
	url := fmt.Sprintf("http://localhost:8080/products/%s", productCode)
	return display(url)
}

func DisplayInvoice(invoiceCode string) error {
	url := fmt.Sprintf("http://localhost:8080/invoices/%s", invoiceCode)
	return display(url)
}

func display(url string) error {
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
