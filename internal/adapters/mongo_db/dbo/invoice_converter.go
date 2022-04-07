package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"log"
)

func ConvertInvoiceToModel(invoice Invoice) model.Invoice {
	return model.NewInvoice(
		invoice.Id,
		invoice.CustomerID,
		invoice.Date,
		convertInvoiceYearMonth(invoice.YearMonth, invoice.Id),
		invoice.ChildrenIds,
		lines(invoice.Lines),
		payment_type.NewPaymentType(invoice.PaymentType),
		invoice.Note,
		invoice.Emailed,
		invoice.Printed,
		invoice.SentToBank,
	)
}

func convertInvoiceYearMonth(yearMonth string, invoiceId string) model.YearMonth {
	ym, err := model.StringToYearMonth(yearMonth)
	if err != nil {
		log.Fatalf("la dada yearMonth '%s' de la factura %s és incorrecte", yearMonth, invoiceId)
	}
	return ym
}

func lines(lines []Line) []model.InvoiceLine {
	var out []model.InvoiceLine
	for _, l := range lines {
		out = append(out, line(l))
	}
	return out
}

func line(line Line) model.InvoiceLine {
	return model.InvoiceLine{
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

func ConvertInvoicesToDbo(invoices []model.Invoice) []interface{} {
	var out []interface{}
	for _, invoice := range invoices {
		out = append(out, ConvertInvoiceToDbo(invoice))
	}
	return out
}

func ConvertInvoiceToDbo(invoice model.Invoice) Invoice {
	var lines []Line
	for _, line := range invoice.Lines() {
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
		Id:          invoice.Id(),
		CustomerID:  invoice.CustomerId(),
		Date:        invoice.Date(),
		YearMonth:   invoice.YearMonth().String(),
		ChildrenIds: invoice.ChildrenIds(),
		Lines:       lines,
		PaymentType: invoice.PaymentType().String(),
		Note:        invoice.Note(),
		Emailed:     invoice.Emailed(),
		Printed:     invoice.Printed(),
		SentToBank:  invoice.SentToBank(),
	}
}
