package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"strings"
)

func ConvertCustomerToModel(customer Customer) model.Customer {
	return model.Customer{
		Id:            customer.Id,
		Active:        customer.Active,
		Children:      children(customer.Children),
		Adults:        adults(customer.Adults),
		InvoiceHolder: holder(customer.InvoiceHolder),
		Note:          customer.Note,
		Language:      customer.Language,
	}
}

func children(children []Child) []model.Child {
	var out []model.Child
	for _, c := range children {
		out = append(out, child(c))
	}
	return out
}

func child(child Child) model.Child {
	return model.Child{
		Id:            child.Id,
		Name:          child.Name,
		Surname:       child.Surname,
		SecondSurname: child.SecondSurname,
		TaxID:         child.TaxID,
		BirthDate:     child.BirthDate,
		Group:         child.Group,
		Note:          child.Note,
		Active:        child.Active,
	}
}

func adults(adults []Adult) []model.Adult {
	var out []model.Adult
	for _, a := range adults {
		out = append(out, adult(a))
	}
	return out
}

func adult(adult Adult) model.Adult {
	return model.Adult{
		Name:             adult.Name,
		Surname:          adult.Surname,
		SecondSurname:    adult.SecondSurname,
		TaxID:            adult.TaxID,
		Role:             adult.Role,
		Address:          address(adult.Address),
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

func address(address Address) model.Address {
	return model.Address{
		Street:  address.Street,
		ZipCode: address.ZipCode,
		City:    address.City,
		State:   address.State,
	}
}

func holder(holder InvoiceHolder) model.InvoiceHolder {
	return model.InvoiceHolder{
		Name:        holder.Name,
		TaxID:       holder.TaxID,
		Address:     address(holder.Address),
		Email:       holder.Email,
		SendEmail:   holder.SendEmail,
		PaymentType: newPaymentType(holder.PaymentType),
		BankAccount: holder.BankAccount,
		IsBusiness:  holder.IsBusiness,
	}
}

func ConvertCustomersToModel(customers []Customer) []model.Customer {
	var out []model.Customer
	for _, customer := range customers {
		out = append(out, ConvertCustomerToModel(customer))
	}
	return out
}

func newPaymentType(value string) payment_type.PaymentType {
	var _values = []string{
		"",
		"BANK_DIRECT_DEBIT",
		"BANK_TRANSFER",
		"VOUCHER",
		"CASH",
		"RECTIFICATION",
	}
	value = strings.ToLower(value)
	for i, val := range _values {
		if strings.ToLower(val) == value {
			return payment_type.PaymentType(i)
		}
	}
	return payment_type.Invalid
}
