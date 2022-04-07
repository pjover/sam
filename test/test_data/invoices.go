package test_data

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
)

var lines = []model.InvoiceLine{
	model.NewInvoiceLine("TST", 1, 11, 0, 1850),
	model.NewInvoiceLine("XXX", 3, 5.5, 0.1, 1850),
	model.NewInvoiceLine("YYY", 1.5, 5, 0, 1850),
}

var InvoiceF100 = model.NewInvoice(
	"F-100",
	148,
	TestDate,
	model.NewYearMonth(2019, 5),
	[]int{1800, 1801},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF101 = model.NewInvoice(
	"F-101",
	148,
	TestDate,
	model.NewYearMonth(2019, 5),
	[]int{1801, 1802},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF102 = model.NewInvoice(
	"F-102",
	149,
	TestDate,
	model.NewYearMonth(2019, 5),
	[]int{1800, 1801, 1802},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)

var InvoiceF103 = model.NewInvoice(
	"F-103",
	149,
	TestDate,
	model.NewYearMonth(2019, 5),
	[]int{1800},
	lines,
	payment_type.BankDirectDebit,
	"Invoice note",
	false,
	false,
	false,
)
