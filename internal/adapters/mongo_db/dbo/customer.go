package dbo

import (
	"time"
)

type Customer struct {
	Id            int           `bson:"_id"`
	Active        bool          `bson:"active"`
	Children      []Child       `bson:"children"`
	Adults        []Adult       `bson:"adults"`
	InvoiceHolder InvoiceHolder `bson:"invoiceHolder"`
	Note          string        `bson:"note"`
	Language      string        `bson:"language"`
	ChangedOn     time.Time     `bson:"changedOn"`
}

type Child struct {
	Id            int       `bson:"code"`
	Name          string    `bson:"name"`
	Surname       string    `bson:"surname"`
	SecondSurname string    `bson:"secondSurname"`
	TaxID         string    `bson:"taxId"`
	BirthDate     time.Time `bson:"birthDate"`
	Group         string    `bson:"group"`
	Note          string    `bson:"note"`
	Active        bool      `bson:"active"`
}

type Adult struct {
	Name             string    `bson:"name"`
	Surname          string    `bson:"surname"`
	SecondSurname    string    `bson:"secondSurname"`
	TaxID            string    `bson:"taxId"`
	Role             string    `bson:"role"`
	Address          Address   `bson:"address"`
	Email            string    `bson:"email"`
	MobilePhone      string    `bson:"mobilePhone"`
	HomePhone        string    `bson:"homePhone"`
	GrandMotherPhone string    `bson:"grandMotherPhone"`
	GrandParentPhone string    `bson:"grandParentPhone"`
	WorkPhone        string    `bson:"workPhone"`
	BirthDate        time.Time `bson:"birthDate"`
	Nationality      string    `bson:"nationality"`
}

type Address struct {
	Street  string `bson:"street"`
	ZipCode string `bson:"zipCode"`
	City    string `bson:"city"`
	State   string `bson:"state"`
}

type InvoiceHolder struct {
	Name        string  `bson:"name"`
	TaxID       string  `bson:"taxId"`
	Address     Address `bson:"address"`
	Email       string  `bson:"email"`
	SendEmail   bool    `bson:"sendEmail"`
	PaymentType string  `bson:"paymentType"`
	Iban        string  `bson:"bankAccount"`
	IsBusiness  bool    `bson:"isBusiness"`
}

func (c Customer) GetId() interface{} {
	return c.Id
}
