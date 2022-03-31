package model

import "fmt"

type Address struct {
	Street  string
	ZipCode string
	City    string
	State   string
}

func (a Address) CompleteAddress() string {
	if a.Street == "" {
		return ""
	}
	return fmt.Sprintf("%s, %s %s, %s", a.Street, a.ZipCode, a.City, a.State)
}
