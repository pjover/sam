package model

import (
	"fmt"
	"strings"
)

func (i Invoice) Code() string {
	url := i.Links.Self.Href
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func (i Invoice) Amount() float64 {
	var amount float64
	for _, line := range i.Lines {
		amount += line.Units * line.ProductPrice
	}
	return amount
}

func (i Invoice) LinesFmt(joinWith string) string {
	var lines []string
	for _, line := range i.Lines {
		lines = append(lines, fmt.Sprintf(
			"%.1f %s (%.2f)",
			line.Units,
			line.ProductID,
			line.Units*line.ProductPrice,
		),
		)
	}
	return strings.Join(lines, joinWith)
}

func (i Invoice) PaymentFmt() string {
	switch i.PaymentType {
	case "BANK_DIRECT_DEBIT":
		return "Rebut"
	case "BANK_TRANSFER":
		return "Tranferència"
	case "CASH":
		return "Efectiu"
	case "VOUCHER":
		return "Xec escoleta"
	default:
		return "Indefinit"
	}
}
