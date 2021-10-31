package adm

import (
	"fmt"
	"net/url"
	"sam/comm"
)

func SearchCustomer(args []string) error {
	text := fmt.Sprint(args)
	params := url.Values{}
	params.Add("text", text)
	_url := fmt.Sprintf("http://localhost:8080/search/customer?%s", params.Encode())
	return comm.PrintGet(_url)
}
