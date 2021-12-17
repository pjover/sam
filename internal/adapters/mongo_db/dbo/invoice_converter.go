package dbo

import "github.com/pjover/sam/internal/core/model"

func ConvertInvoice(invoice Invoice) model.Invoice {
	return model.Invoice{
		Id:            invoice.Id,
		CustomerID:    invoice.CustomerID,
		Date:          invoice.Date,
		YearMonth:     invoice.YearMonth,
		ChildrenCodes: invoice.ChildrenCodes,
		Lines:         lines(invoice.Lines),
		PaymentType:   invoice.PaymentType,
		Note:          invoice.Note,
		Emailed:       invoice.Emailed,
		Printed:       invoice.Printed,
		SentToBank:    invoice.SentToBank,
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
		ProductID:     line.ProductID,
		Units:         line.Units,
		ProductPrice:  line.ProductPrice,
		TaxPercentage: line.TaxPercentage,
		ChildCode:     line.ChildCode,
	}
}
