package model

import "time"

type Child struct {
	Id            int
	Name          string
	Surname       string
	SecondSurname string
	TaxID         string
	BirthDate     time.Time
	Group         string
	Note          string
	Active        bool
}

type Adult struct {
	Name             string
	Surname          string
	SecondSurname    string
	TaxID            string
	Role             string
	Address          Address
	Email            string
	MobilePhone      string
	HomePhone        string
	GrandMotherPhone string
	GrandParentPhone string
	WorkPhone        string
	BirthDate        time.Time
	Nationality      string
}

type Address struct {
	Street  string
	ZipCode string
	City    string
	State   string
}

type InvoiceHolder struct {
	Name        string
	TaxID       string
	Address     Address
	Email       string
	SendEmail   bool
	PaymentType string
	BankAccount string
	IsBusiness  bool
}

type Customer struct {
	Id            int
	Active        bool
	Children      []Child
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          interface{}
	Language      string
}
