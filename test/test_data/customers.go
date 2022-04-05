package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var Child1480 = model.NewChild(
	1480,
	"Laura",
	"Llull",
	"Bibiloni",
	model.NewTaxIdOrEmpty(""),
	TestDate,
	group_type.EI_1,
	"Note child 1480",
	true,
)

var Child1481 = model.NewChild(
	1481,
	"Aina",
	"Llull",
	"Bibiloni",
	model.NewTaxIdOrEmpty("60235657Z"),
	TestDate,
	group_type.EI_1,
	"Note child 1481",
	true,
)

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

var InvoiceHolder148 = model.NewInvoiceHolder(
	"Cara Santamaria Novella",
	model.NewTaxIdOrEmpty("36361882D"),
	model.NewAddress(
		"Carrer Ucraïna 2022, 1st",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"cara@sgu.org",
	false,
	payment_type.BankDirectDebit,
	model.NewIbanOrEmpty("ES2830668859978258529057"),
	false,
)

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

var Child1490 = model.NewChild(
	1490,
	"Antònia",
	"Petit",
	"Petita",
	model.NewTaxIdOrEmpty("81620787C"),
	TestDate,
	group_type.EI_2,
	"Note child 1490",
	true,
)

var Child1491 = model.NewChild(
	1491,
	"Antoni",
	"Petit",
	"Petita",
	model.NewTaxIdOrEmpty("51389353Q"),
	TestDate,
	group_type.EI_3,
	"Note child 1491",
	true,
)

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

var InvoiceHolder149 = model.NewInvoiceHolder(
	"Nom empresa",
	model.NewTaxIdOrEmpty("37866397W"),
	model.NewAddress(
		"Address first line",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"email@gmail.com",
	false,
	payment_type.BankTransfer,
	model.NewIbanOrEmpty("ES2830668859978258529057"),
	true,
)

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
