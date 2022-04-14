package model

import (
	"github.com/pjover/sam/internal/domain/model/adult_role"
	"github.com/pjover/sam/internal/domain/model/group_type"
	"github.com/pjover/sam/internal/domain/model/language"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"time"
)

var TestDate = time.Date(2019, 5, 25, 0, 0, 0, 0, time.UTC)

var TestAdultMother148 = NewAdult(
	"Cara",
	"Santamaria",
	"Novella",
	NewTaxIdOrEmpty("36361882D"),
	adult_role.Mother,
	NewAddress(
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
	NewNationalityOrEmpty("US"),
)

var TestAdultFather148 = NewAdult(
	"Bob",
	"Novella",
	"Sagan",
	NewTaxIdOrEmpty("71032204Q"),
	adult_role.Father,
	NewAddress(
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
	NewNationalityOrEmpty("UK"),
)

var TestAdultMother149 = NewAdult(
	"Joana",
	"Petita",
	"Puig",
	NewTaxIdOrEmpty("80587890F"),
	adult_role.Mother,
	NewAddress(
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
	NewNationalityOrEmpty("ES"),
)

var TestAdultFather149 = NewAdult(
	"Joan",
	"Petit",
	"Galatzó",
	NewTaxIdOrEmpty("91071996T"),
	adult_role.Father,
	NewAddress(
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
	NewNationalityOrEmpty("ES"),
)

var TestChild1480 = NewChild(
	1480,
	"Laura",
	"Llull",
	"Bibiloni",
	NewTaxIdOrEmpty(""),
	TestDate,
	group_type.Ei1,
	"Note child 1480",
	true,
)

var TestChild1481 = NewChild(
	1481,
	"Aina",
	"Llull",
	"Bibiloni",
	NewTaxIdOrEmpty("60235657Z"),
	TestDate,
	group_type.Ei1,
	"Note child 1481",
	true,
)

var TestChild1490 = NewChild(
	1490,
	"Antònia",
	"Petit",
	"Petita",
	NewTaxIdOrEmpty("81620787C"),
	TestDate,
	group_type.Ei2,
	"Note child 1490",
	true,
)

var TestChild1491 = NewChild(
	1491,
	"Antoni",
	"Petit",
	"Petita",
	NewTaxIdOrEmpty("51389353Q"),
	TestDate,
	group_type.Ei3,
	"Note child 1491",
	true,
)

var TestCustomer148 = NewCustomer(
	148,
	true,
	[]Child{
		TestChild1480,
		TestChild1481,
	},
	[]Adult{
		TestAdultMother148,
		TestAdultFather148,
	},
	TestInvoiceHolder148,
	"Nota del client 148",
	language.Catalan,
	TestDate,
)

var TestCustomer149 = NewCustomer(
	149,
	true,
	[]Child{
		TestChild1490,
		TestChild1491,
	},
	[]Adult{
		TestAdultMother149,
		TestAdultFather149,
	},
	TestInvoiceHolder149,
	"Nota del client 149",
	language.Catalan,
	TestDate,
)

var TestInvoiceHolder148 = NewInvoiceHolder(
	"Cara Santamaria Novella",
	NewTaxIdOrEmpty("36361882D"),
	NewAddress(
		"Carrer Ucraïna 2022, 1st",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"cara@sgu.org",
	false,
	payment_type.BankDirectDebit,
	NewIbanOrEmpty("ES2830668859978258529057"),
	false,
)

var TestInvoiceHolder149 = NewInvoiceHolder(
	"Nom empresa",
	NewTaxIdOrEmpty("37866397W"),
	NewAddress(
		"Address first line",
		"07007",
		"Palma",
		"Illes Balears",
	),
	"email@gmail.com",
	false,
	payment_type.BankTransfer,
	NewIbanOrEmpty("ES2830668859978258529057"),
	true,
)

var lines = []InvoiceLine{
	NewInvoiceLine("TST", 1, 11, 0, 1850),
	NewInvoiceLine("XXX", 3, 5.5, 0.1, 1850),
	NewInvoiceLine("YYY", 1.5, 5, 0, 1850),
}

var InvoiceF100 = NewInvoice(
	"F-100",
	148,
	TestDate,
	NewYearMonth(2019, 5),
	[]int{1800, 1801},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF101 = NewInvoice(
	"F-101",
	148,
	TestDate,
	NewYearMonth(2019, 5),
	[]int{1801, 1802},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF102 = NewInvoice(
	"F-102",
	149,
	TestDate,
	NewYearMonth(2019, 5),
	[]int{1800, 1801, 1802},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF103 = NewInvoice(
	"F-103",
	149,
	TestDate,
	NewYearMonth(2019, 5),
	[]int{1800},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var ProductTST, _ = NewProduct(
	"TST",
	"Test product",
	"TstProduct",
	10.9,
	0.0,
	false,
)

var ProductXXX, _ = NewProduct(
	"XXX",
	"XXX product",
	"XxxProduct",
	9.1,
	0.0,
	false,
)

var ProductYYY, _ = NewProduct(
	"YYY",
	"YYY product",
	"YyyProduct",
	5,
	0.1,
	false,
)
