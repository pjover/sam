package model

import (
	"fmt"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"sort"
	"strings"
	"time"
)

type Invoice struct {
	id          string
	customerId  int
	date        time.Time
	yearMonth   YearMonth
	childrenIds []int
	lines       []InvoiceLine
	paymentType payment_type.PaymentType
	note        string
	emailed     bool
	printed     bool
	sentToBank  bool
}

func NewInvoice(
	id string,
	customerId int,
	date time.Time,
	yearMonth YearMonth,
	childrenIds []int,
	lines []InvoiceLine,
	paymentType payment_type.PaymentType,
	note string,
	emailed bool,
	printed bool,
	sentToBank bool,
) Invoice {
	return Invoice{
		id:          id,
		customerId:  customerId,
		date:        date,
		yearMonth:   yearMonth,
		childrenIds: childrenIds,
		lines:       lines,
		paymentType: paymentType,
		note:        note,
		emailed:     emailed,
		printed:     printed,
		sentToBank:  sentToBank,
	}
}

func (i Invoice) Id() string {
	return i.id
}

func (i Invoice) CustomerId() int {
	return i.customerId
}

func (i Invoice) Date() time.Time {
	return i.date
}

func (i Invoice) YearMonth() YearMonth {
	return i.yearMonth
}

func (i Invoice) ChildrenIds() []int {
	return i.childrenIds
}

func (i Invoice) Lines() []InvoiceLine {
	return i.lines
}

func (i Invoice) PaymentType() payment_type.PaymentType {
	return i.paymentType
}

func (i Invoice) Note() string {
	return i.note
}

func (i Invoice) Emailed() bool {
	return i.emailed
}

func (i Invoice) Printed() bool {
	return i.printed
}

func (i Invoice) SentToBank() bool {
	return i.sentToBank
}

func (i Invoice) String() string {
	return fmt.Sprintf("%d  %s  %s  %7.2f  %s  %s", i.customerId, i.id, i.yearMonth.String(), i.Amount(), i.paymentType.Format(), i.LinesFmt(","))

}

func (i Invoice) DateFmt() string {
	return i.date.Format(domain.YearMonthDayLayout)
}

func (i Invoice) Amount() float64 {
	var amount float64
	for _, line := range i.lines {
		gross := line.Units() * line.ProductPrice()
		tax := gross * line.TaxPercentage()
		amount += gross + tax
	}
	return amount
}

func (i Invoice) LinesFmt(joinWith string) string {
	var lines []string
	for _, line := range i.lines {
		lines = append(lines, line.Format())
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i] < lines[j]
	})
	return strings.Join(lines, joinWith)
}
