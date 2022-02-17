package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"strings"
)

func ConvertInvoiceToModel(invoice Invoice) model.Invoice {
	return model.Invoice{
		Id:          invoice.Id,
		CustomerId:  invoice.CustomerID,
		Date:        invoice.Date,
		YearMonth:   invoice.YearMonth,
		ChildrenIds: invoice.ChildrenIds,
		Lines:       lines(invoice.Lines),
		PaymentType: convertToPaymentType(invoice.PaymentType),
		Note:        invoice.Note,
		Emailed:     invoice.Emailed,
		Printed:     invoice.Printed,
		SentToBank:  invoice.SentToBank,
	}
}

func lines(lines []Line) []model.Line {
	var out []model.Line
	for _, l := range lines {
		out = append(out, line(l))
	}
	return out
}

func line(line Line) model.Line {
	return model.Line{
		ProductId:     line.ProductID,
		Units:         Decimal128ToFloat64(line.Units),
		ProductPrice:  Decimal128ToFloat64(line.ProductPrice),
		TaxPercentage: Decimal128ToFloat64(line.TaxPercentage),
		ChildId:       line.ChildId,
	}
}

func ConvertInvoicesToModel(invoices []Invoice) []model.Invoice {
	var out []model.Invoice
	for _, invoice := range invoices {
		out = append(out, ConvertInvoiceToModel(invoice))
	}
	return out
}

var values = []string{
	"BANK_DIRECT_DEBIT",
	"BANK_TRANSFER",
	"VOUCHER",
	"CASH",
	"RECTIFICATION",
}

func convertToPaymentType(value string) payment_type.PaymentType {
	value = strings.ToLower(value)
	for i, val := range values {
		if strings.ToLower(val) == value {
			return payment_type.PaymentType(i)
		}
	}
	return payment_type.Invalid
}
