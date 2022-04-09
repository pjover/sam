package model

import (
	"errors"
	"fmt"
	"strconv"
)

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

func (a Address) validate() error {
	if a.street == "" {
		return a.validateEmpty()
	}

	if a.zipCode == "" {
		return errors.New("el codi postal (ZipCode) no pot estar buit")
	}
	if len(a.zipCode) != 5 {
		return errors.New("el codi postal (ZipCode) ha de tenir 5 números")
	}
	_, err := strconv.Atoi(a.zipCode)
	if err != nil {
		return errors.New("el codi postal (ZipCode) només pot tenir números")
	}

	if a.city == "" {
		return errors.New("la ciutat (City) no pot estar buida")
	}

	if a.state == "" {
		return errors.New("el estat o regió (State) no pot estar buit")
	}
	return nil
}

func (a Address) validateEmpty() error {
	err := errors.New("si el carrer (Street) és buit, la resta de camps han d'esser buits")
	if a.zipCode != "" {
		return err
	}
	if a.city != "" {
		return err
	}
	if a.state != "" {
		return err
	}
	return nil
}
