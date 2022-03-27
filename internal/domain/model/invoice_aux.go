package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"sort"
	"strings"
)

func (i Invoice) String() string {
	return fmt.Sprintf("%d  %s  %s  % 7.2f  %s  %s", i.CustomerId, i.Id, i.YearMonth.String(), i.Amount(), i.PaymentType.String(), i.LinesFmt(","))

}

func (i Invoice) DateFmt() string {
	return i.Date.Format(domain.YearMonthDayLayout)
}

func (i Invoice) Amount() float64 {
	var amount float64
	for _, line := range i.Lines {
		gross := line.Units * line.ProductPrice
		tax := gross * line.TaxPercentage
		amount += gross + tax
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
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	return strings.Join(lines, joinWith)
}
