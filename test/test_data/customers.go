package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var Child1850 = model.Child{
	Id:            1850,
	Name:          "Laura",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	BirthDate:     TestDate,
	Group:         group_type.EI_1,
	Active:        true,
}

var Child1851 = model.Child{
	Id:            1851,
	Name:          "Aina",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	TaxID:         model.NewTaxIdOrEmpty("60235657Z"),
	BirthDate:     TestDate,
	Group:         group_type.EI_1,
	Active:        true,
}

var AdultMother = model.NewAdult(
	"Cara",
	"Santamaria",
	"Novella",
	model.NewTaxIdOrEmpty("36361882D"),
	adult_role.Mother,
	model.NewAddress(
		"Carrer Ucraïna 2022, 1st",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"cara@sgu.org",
	"654321098",
	"987654321",
	"685698789",
	"658785478",
	"987525444",
	TestDate,
	model.NewNationalityOrEmpty("US"),
)

var AdultFather = model.NewAdult(
	"Bob",
	"Novella",
	"Sagan",
	model.NewTaxIdOrEmpty("71032204Q"),
	adult_role.Father,
	model.NewAddress(
		"Carrer Ucraïna 2022, 1st",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"bob@sgu.org",
	"654321097",
	"987654322",
	"685698788",
	"658785477",
	"987525446",
	TestDate,
	model.NewNationalityOrEmpty("UK"),
)

var InvoiceHolder148 = model.InvoiceHolder{
	Name:  "Cara Santamaria Novella",
	TaxID: model.NewTaxIdOrEmpty("36361882D"),
	Address: model.NewAddress(
		"Carrer Ucraïna 2022, 1st",
		"07007",
		"Palma",
		"Illes Balears",
	),
	Email:       "cara@sgu.org",
	PaymentType: payment_type.BankDirectDebit,
	Iban:        model.NewIbanOrEmpty("ES2830668859978258529057"),
}

var InvoiceHolder149 = model.InvoiceHolder{
	Name:  "Nom empresa",
	TaxID: model.NewTaxIdOrEmpty("37866397W"),
	Address: model.NewAddress(
		"Address first line",
		"07007",
		"Palma",
		"Illes Balears",
	),
	Email:       "email@gmail.com",
	PaymentType: payment_type.BankTransfer,
	Iban:        model.NewIbanOrEmpty("ES2830668859978258529057"),
	IsBusiness:  true,
}

var Customer148 = model.Customer{
	Id:     148,
	Active: true,
	Children: []model.Child{
		Child1850,
		Child1851,
	},
	Adults: []model.Adult{
		AdultMother,
		AdultFather,
	},
	InvoiceHolder: InvoiceHolder148,
	Note:          "Nota del client",
	Language:      language.Catalan,
}

var Customer149 = model.Customer{
	Id:     149,
	Active: true,
	Children: []model.Child{
		Child1850,
		Child1851,
	},
	Adults: []model.Adult{
		AdultMother,
		AdultFather,
	},
	InvoiceHolder: InvoiceHolder149,
	Note:          "Nota del client",
	Language:      language.Catalan,
}
