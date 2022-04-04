package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

func ConvertCustomerToModel(customer Customer) model.Customer {
	return model.Customer{
		Id:            customer.Id,
		Active:        customer.Active,
		Children:      children(customer.Children),
		Adults:        adults(customer.Adults),
		InvoiceHolder: holderToModel(customer.InvoiceHolder),
		Note:          customer.Note,
		Language:      language.NewLanguage(customer.Language),
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
		TaxID:         model.NewTaxIdOrEmpty(child.TaxID),
		BirthDate:     child.BirthDate,
		Group:         group_type.NewGroupType(child.Group),
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
		TaxID:            model.NewTaxIdOrEmpty(adult.TaxID),
		Role:             adult_role.NewAdultRole(adult.Role),
		Address:          address(adult.Address),
		Email:            adult.Email,
		MobilePhone:      adult.MobilePhone,
		HomePhone:        adult.HomePhone,
		GrandMotherPhone: adult.GrandMotherPhone,
		GrandParentPhone: adult.GrandParentPhone,
		WorkPhone:        adult.WorkPhone,
		BirthDate:        adult.BirthDate,
		Nationality:      model.NewNationalityOrEmpty(adult.Nationality),
	}
}

func address(address Address) model.Address {
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
		Address:     address(holder.Address),
		Email:       holder.Email,
		SendEmail:   holder.SendEmail,
		PaymentType: payment_type.NewPaymentType(holder.PaymentType),
		Iban:        model.NewIbanOrEmpty(holder.Iban),
		IsBusiness:  holder.IsBusiness,
	}
}

func CustomerToModel(customers []Customer) []model.Customer {
	var out []model.Customer
	for _, customer := range customers {
		out = append(out, ConvertCustomerToModel(customer))
	}
	return out
}

func ConvertCustomerToDbo(customer model.Customer) Customer {
	return Customer{
		Id:            customer.Id,
		Active:        customer.Active,
		Children:      childrenToDbo(customer.Children),
		Adults:        adultsToDbo(customer.Adults),
		InvoiceHolder: invoiceHolderToDbo(customer.InvoiceHolder),
		Note:          customer.Note,
		Language:      customer.Language.String(),
		ChangedOn:     customer.ChangedOn,
	}
}

func childrenToDbo(children []model.Child) []Child {
	var _children []Child
	for _, child := range children {
		_children = append(_children, childToDbo(child))
	}
	return _children
}

func childToDbo(child model.Child) Child {
	return Child{
		Id:            child.Id,
		Name:          child.Name,
		Surname:       child.Surname,
		SecondSurname: child.SecondSurname,
		TaxID:         child.TaxID.String(),
		BirthDate:     child.BirthDate,
		Group:         child.Group.String(),
		Note:          child.Note,
		Active:        child.Active,
	}
}

func adultsToDbo(adults []model.Adult) []Adult {
	var _adults []Adult
	for _, adult := range adults {
		_adults = append(_adults, adultToDbo(adult))
	}
	return _adults
}

func adultToDbo(adult model.Adult) Adult {
	return Adult{
		Name:             adult.Name,
		Surname:          adult.Surname,
		SecondSurname:    adult.SecondSurname,
		TaxID:            adult.TaxID.String(),
		Role:             adult.Role.String(),
		Address:          addressToDbo(adult.Address),
		Email:            adult.Email,
		MobilePhone:      adult.MobilePhone,
		HomePhone:        adult.HomePhone,
		GrandMotherPhone: adult.GrandMotherPhone,
		GrandParentPhone: adult.GrandParentPhone,
		WorkPhone:        adult.WorkPhone,
		BirthDate:        adult.BirthDate,
		Nationality:      adult.Nationality.String(),
	}
}

func invoiceHolderToDbo(invoiceHolder model.InvoiceHolder) InvoiceHolder {
	return InvoiceHolder{
		Name:        invoiceHolder.Name,
		TaxID:       invoiceHolder.TaxID.String(),
		Address:     addressToDbo(invoiceHolder.Address),
		Email:       invoiceHolder.Email,
		SendEmail:   invoiceHolder.SendEmail,
		PaymentType: invoiceHolder.PaymentType.String(),
		Iban:        invoiceHolder.Iban.String(),
		IsBusiness:  invoiceHolder.IsBusiness,
	}
}

func addressToDbo(address model.Address) Address {
	return Address{
		Street:  address.Street(),
		ZipCode: address.ZipCode(),
		City:    address.City(),
		State:   address.State(),
	}
}
