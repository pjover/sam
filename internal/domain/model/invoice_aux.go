package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"strings"
)

func (i Invoice) String() string {
	return fmt.Sprintf("%d  %s  %s  % 7.2f  %s  %s", i.CustomerId, i.Id, i.YearMonth, i.Amount(), i.PaymentType.String(), i.LinesFmt(","))

}

func (i Invoice) DateFmt() string {
	return i.Date.Format(domain.YearMonthDayLayout)
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
			line.ProductId,
			line.Units*line.ProductPrice,
		),
		)
	}
	return strings.Join(lines, joinWith)
}
