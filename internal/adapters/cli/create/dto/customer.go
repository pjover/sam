package dto

import (
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"log"
	"time"
)

type TransientCustomer struct {
	Children      []TransientChild
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          string
	Language      string
}

type Customer struct {
	Id            int
	Active        bool
	Children      []Child
	Adults        []Adult
	InvoiceHolder InvoiceHolder
	Note          string
	Language      string
	ChangedOn     time.Time
}

type TransientChild struct {
	Name          string
	Surname       string
	SecondSurname string
	TaxID         string
	BirthDate     string
	Group         string
	Note          string
}

type Child struct {
	Id            int
	Name          string
	Surname       string
	SecondSurname string
	TaxID         string
	BirthDate     string
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
	BirthDate        string
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
	Iban        string
	IsBusiness  bool
}

func TransientCustomerToModel(customer TransientCustomer) model.TransientCustomer {
	return model.TransientCustomer{
		transientChildrenToModel(customer.Children),
		adultsToModel(customer.Adults),
		holderToModel(customer.InvoiceHolder),
		customer.Note,
		language.NewLanguage(customer.Language),
	}
}

func transientChildrenToModel(children []TransientChild) []model.TransientChild {
	var out []model.TransientChild
	for _, c := range children {
		out = append(out, childToModel(c))
	}
	return out
}

func childToModel(child TransientChild) model.TransientChild {
	return model.TransientChild{
		Name:          child.Name,
		Surname:       child.Surname,
		SecondSurname: child.SecondSurname,
		TaxID:         model.NewTaxIdOrEmpty(child.TaxID),
		BirthDate:     stringToTime(child.BirthDate),
		Group:         group_type.NewGroupType(child.Group),
		Note:          child.Note,
	}
}

func stringToTime(date string) time.Time {
	if date == "" {
		return time.Time{}
	}
	tm, err := time.Parse(domain.YearMonthDayLayout, date)
	if err != nil {
		log.Fatalf("error de format en la data '%s', ha de tenir el format 2022-03-29", date)
	}
	return tm
}

func adultsToModel(adults []Adult) []model.Adult {
	var out []model.Adult
	for _, a := range adults {
		out = append(out, adultToModel(a))
	}
	return out
}

func adultToModel(adult Adult) model.Adult {
	return model.NewAdult(
		adult.Name,
		adult.Surname,
		adult.SecondSurname,
		model.NewTaxIdOrEmpty(adult.TaxID),
		adult_role.NewAdultRole(adult.Role),
		addressToModel(adult.Address),
		adult.Email,
		adult.MobilePhone,
		adult.HomePhone,
		adult.GrandMotherPhone,
		adult.GrandParentPhone,
		adult.WorkPhone,
		stringToTime(adult.BirthDate),
		model.NewNationalityOrEmpty(adult.Nationality),
	)
}

func addressToModel(address Address) model.Address {
	return model.NewAddress(
		address.Street,
		address.ZipCode,
		address.City,
		address.State,
	)
}

func holderToModel(holder InvoiceHolder) model.InvoiceHolder {
	return model.InvoiceHolder{
		Name:        holder.Name,
		TaxID:       model.NewTaxIdOrEmpty(holder.TaxID),
		Address:     addressToModel(holder.Address),
		Email:       holder.Email,
		SendEmail:   holder.SendEmail,
		PaymentType: payment_type.NewPaymentType(holder.PaymentType),
		Iban:        model.NewIbanOrEmpty(holder.Iban),
		IsBusiness:  holder.IsBusiness,
	}
}
