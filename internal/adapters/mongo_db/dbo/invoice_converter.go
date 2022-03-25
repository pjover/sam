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

func ConvertInvoicesToDbo(invoices []model.Invoice) []interface{} {
	var out []interface{}
	for _, invoice := range invoices {
		out = append(out, ConvertInvoiceToDbo(invoice))
	}
	return out
}

func ConvertInvoiceToDbo(invoice model.Invoice) Invoice {
	var lines []Line
	for _, line := range invoice.Lines {
		_line := Line{
			ProductID:     line.ProductId,
			Units:         Float64ToDecimal128(line.Units),
			ProductPrice:  Float64ToDecimal128(line.ProductPrice),
			TaxPercentage: Float64ToDecimal128(line.TaxPercentage),
			ChildId:       line.ChildId,
		}
		lines = append(lines, _line)
	}
	return Invoice{
		Id:          invoice.Id,
		CustomerID:  invoice.CustomerId,
		Date:        invoice.Date,
		YearMonth:   invoice.YearMonth,
		ChildrenIds: invoice.ChildrenIds,
		Lines:       lines,
		PaymentType: PaymentTypes[invoice.PaymentType],
		Note:        invoice.Note,
		Emailed:     invoice.Emailed,
		Printed:     invoice.Printed,
		SentToBank:  invoice.SentToBank,
	}
}

var PaymentTypes = []string{
	"",
	"BANK_DIRECT_DEBIT",
	"BANK_TRANSFER",
	"VOUCHER",
	"CASH",
	"RECTIFICATION",
}
