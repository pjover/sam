package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
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
		Language:      model.NewLanguage(customer.Language),
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
		Group:         newGroupType(child.Group),
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
		Role:             newAdultRole(adult.Role),
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

var adultRoleValues = []string{
	"",
	"MOTHER",
	"FATHER",
	"TUTOR",
}

func newAdultRole(value string) adult_role.AdultRole {
	value = strings.ToLower(value)
	for i, val := range adultRoleValues {
		if strings.ToLower(val) == value {
			return adult_role.AdultRole(i)
		}
	}
	return adult_role.Invalid
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

var paymentTypeValues = []string{
	"",
	"BANK_DIRECT_DEBIT",
	"BANK_TRANSFER",
	"VOUCHER",
	"CASH",
	"RECTIFICATION",
}

func newPaymentType(value string) payment_type.PaymentType {
	value = strings.ToLower(value)
	for i, val := range paymentTypeValues {
		if strings.ToLower(val) == value {
			return payment_type.PaymentType(i)
		}
	}
	return payment_type.Invalid
}

var groupValues = []string{
	"UNDEFINED",
	"EI_1",
	"EI_2",
	"EI_3",
}

func newGroupType(value string) group_type.GroupType {

	value = strings.ToLower(value)
	for i, val := range groupValues {
		if strings.ToLower(val) == value {
			return group_type.GroupType(i)
		}
	}
	return group_type.Undefined
}

func ConvertCustomerToDbo(customer model.Customer) Customer {
	return Customer{
		Id:            customer.Id,
		Active:        customer.Active,
		Children:      convertChildrenToDbo(customer.Children),
		Adults:        convertAdultsToDbo(customer.Adults),
		InvoiceHolder: convertInvoiceHolderToDbo(customer.InvoiceHolder),
		Note:          customer.Note,
		Language:      customer.Language.String(),
		ChangedOn:     customer.ChangedOn,
	}
}

func convertChildrenToDbo(children []model.Child) []Child {
	var _children []Child
	for _, child := range children {
		_children = append(_children, convertChildToDbo(child))
	}
	return _children
}

func convertChildToDbo(child model.Child) Child {
	return Child{
		Id:            child.Id,
		Name:          child.Name,
		Surname:       child.Surname,
		SecondSurname: child.SecondSurname,
		TaxID:         child.TaxID,
		BirthDate:     child.BirthDate,
		Group:         groupValues[child.Group],
		Note:          child.Note,
		Active:        child.Active,
	}
}

func convertAdultsToDbo(adults []model.Adult) []Adult {
	var _adults []Adult
	for _, adult := range adults {
		_adults = append(_adults, convertAdultToDbo(adult))
	}
	return _adults
}

func convertAdultToDbo(adult model.Adult) Adult {
	return Adult{
		Name:             adult.Name,
		Surname:          adult.Surname,
		SecondSurname:    adult.SecondSurname,
		TaxID:            adult.TaxID,
		Role:             adultRoleValues[adult.Role],
		Address:          convertAddressToDbo(adult.Address),
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

func convertInvoiceHolderToDbo(invoiceHolder model.InvoiceHolder) InvoiceHolder {
	return InvoiceHolder{
		Name:        invoiceHolder.Name,
		TaxID:       invoiceHolder.TaxID,
		Address:     convertAddressToDbo(invoiceHolder.Address),
		Email:       invoiceHolder.Email,
		SendEmail:   invoiceHolder.SendEmail,
		PaymentType: paymentTypeValues[invoiceHolder.PaymentType],
		BankAccount: invoiceHolder.BankAccount,
		IsBusiness:  invoiceHolder.IsBusiness,
	}
}

func convertAddressToDbo(address model.Address) Address {
	return Address{
		Street:  address.Street,
		ZipCode: address.ZipCode,
		City:    address.City,
		State:   address.State,
	}
}
