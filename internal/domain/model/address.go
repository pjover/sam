package model

import "fmt"

type Address struct {
	street  string
	zipCode string
	city    string
	state   string
}

func NewAddress(street string, zipCode string, city string, state string) Address {
	return Address{
		street:  street,
		zipCode: zipCode,
		city:    city,
		state:   state,
	}
}

func (a Address) Street() string {
	return a.street
}

func (a Address) ZipCode() string {
	return a.zipCode
}

func (a Address) City() string {
	return a.city
}

func (a Address) State() string {
	return a.state
}

func (a Address) CompleteAddress() string {
	if a.street == "" {
		return ""
	}
	return fmt.Sprintf("%s, %s %s, %s", a.street, a.zipCode, a.city, a.state)
}
