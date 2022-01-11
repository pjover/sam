package model

import (
	"fmt"
	"github.com/pjover/sam/internal/core"
	"strings"
)

func (i Invoice) String() string {
	return fmt.Sprintf("%d  %s  %s  % 7.2f  %s  %s", i.CustomerID, i.Code, i.YearMonth, i.Amount(), i.PaymentFmt(), i.LinesFmt(","))

}

func (i Invoice) DateFmt() string {
	return i.Date.Format(core.YearMonthDayLayout)
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
