package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var Child1480 = model.Child{
	Id:            1480,
	Name:          "Laura",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	BirthDate:     TestDate,
	Group:         group_type.EI_1,
	Active:        true,
}

var Child1481 = model.Child{
	Id:            1481,
	Name:          "Aina",
	Surname:       "Llull",
	SecondSurname: "Bibiloni",
	TaxID:         model.NewTaxIdOrEmpty("60235657Z"),
	BirthDate:     TestDate,
	Group:         group_type.EI_1,
	Active:        true,
}

var AdultMother148 = model.NewAdult(
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

var AdultFather148 = model.NewAdult(
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

var Customer148 = model.NewCustomer(
	148,
	true,
	[]model.Child{
		Child1480,
		Child1481,
	},
	[]model.Adult{
		AdultMother148,
		AdultFather148,
	},
	InvoiceHolder148,
	"Nota del client 148",
	language.Catalan,
	TestDate,
)

var Child1490 = model.Child{
	Id:            1480,
	Name:          "Antònia",
	Surname:       "Petit",
	SecondSurname: "Petita",
	TaxID:         model.NewTaxIdOrEmpty("81620787C"),
	BirthDate:     TestDate,
	Group:         group_type.EI_2,
	Active:        true,
}

var Child1491 = model.Child{
	Id:            1481,
	Name:          "Antoni",
	Surname:       "Petit",
	SecondSurname: "Petita",
	TaxID:         model.NewTaxIdOrEmpty("51389353Q"),
	BirthDate:     TestDate,
	Group:         group_type.EI_3,
	Active:        true,
}

var AdultMother149 = model.NewAdult(
	"Joana",
	"Petita",
	"Puig",
	model.NewTaxIdOrEmpty("80587890F"),
	adult_role.Mother,
	model.NewAddress(
		"Carrer de sa Tanca 2, 1er",
		"07192",
		"Estellencs",
		"Illes Balears",
	),
	"joana@cameva.org",
	"654521098",
	"987674321",
	"695698789",
	"657785478",
	"987524444",
	TestDate,
	model.NewNationalityOrEmpty("ES"),
)

var AdultFather149 = model.NewAdult(
	"Joan",
	"Petit",
	"Galatzó",
	model.NewTaxIdOrEmpty("91071996T"),
	adult_role.Father,
	model.NewAddress(
		"Carrer de sa Tanca 2, 1er",
		"07192",
		"Estellencs",
		"Illes Balears",
	),
	"joan@cameva.org",
	"654321099",
	"987654329",
	"685698789",
	"658785479",
	"987525449",
	TestDate,
	model.NewNationalityOrEmpty("ES"),
)

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

var Customer149 = model.NewCustomer(
	149,
	true,
	[]model.Child{
		Child1490,
		Child1491,
	},
	[]model.Adult{
		AdultMother149,
		AdultFather149,
	},
	InvoiceHolder149,
	"Nota del client 149",
	language.Catalan,
	TestDate,
)
