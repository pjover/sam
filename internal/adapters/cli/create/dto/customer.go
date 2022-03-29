package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"time"
)

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

func CustomerToModel(customer Customer) model.Customer {
	return model.Customer{
		Id:            customer.Id,
		Active:        customer.Active,
		Children:      childrenToModel(customer.Children),
		Adults:        adultsToModel(customer.Adults),
		InvoiceHolder: holderToModel(customer.InvoiceHolder),
		Note:          customer.Note,
		Language:      model.NewLanguage(customer.Language),
	}
}

func childrenToModel(children []Child) []model.Child {
	var out []model.Child
	for _, c := range children {
		out = append(out, childToModel(c))
	}
	return out
}

func childToModel(child Child) model.Child {
	return model.Child{
		Id:            child.Id,
		Name:          child.Name,
		Surname:       child.Surname,
		SecondSurname: child.SecondSurname,
		TaxID:         child.TaxID,
		BirthDate:     child.BirthDate,
		Group:         group_type.NewGroupType(child.Group),
		Note:          child.Note,
		Active:        child.Active,
	}
}

func adultsToModel(adults []Adult) []model.Adult {
	var out []model.Adult
	for _, a := range adults {
		out = append(out, adultToModel(a))
	}
	return out
}

func adultToModel(adult Adult) model.Adult {
	return model.Adult{
		Name:             adult.Name,
		Surname:          adult.Surname,
		SecondSurname:    adult.SecondSurname,
		TaxID:            adult.TaxID,
		Role:             adult_role.NewAdultRole(adult.Role),
		Address:          addressToModel(adult.Address),
		Email:            adult.Email,
		MobilePhone:      adult.MobilePhone,
		HomePhone:        adult.HomePhone,
		GrandMotherPhone: adult.GrandMotherPhone,
		GrandParentPhone: adult.GrandParentPhone,
		WorkPhone:        adult.WorkPhone,
		BirthDate:        adult.BirthDate,
		Nationality:      adult.Nationality,
	}
}

func addressToModel(address Address) model.Address {
	return model.Address{
		Street:  address.Street,
		ZipCode: address.ZipCode,
		City:    address.City,
		State:   address.State,
	}
}

func holderToModel(holder InvoiceHolder) model.InvoiceHolder {
	return model.InvoiceHolder{
		Name:        holder.Name,
		TaxID:       holder.TaxID,
		Address:     addressToModel(holder.Address),
		Email:       holder.Email,
		SendEmail:   holder.SendEmail,
		PaymentType: payment_type.NewPaymentType(holder.PaymentType),
		BankAccount: holder.BankAccount,
		IsBusiness:  holder.IsBusiness,
	}
}
