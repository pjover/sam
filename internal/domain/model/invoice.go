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
		lineAmount := gross * (1 + line.TaxPercentage())
		amount += lineAmount
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

// SendToBank sets sentToBank to true
func (i Invoice) SendToBank() Invoice {
	i.sentToBank = true
	return i
}

type TransientInvoice struct {
	IsRectification bool
	CustomerId      int
	Date            time.Time
	YearMonth       YearMonth
	ChildrenIds     []int
	Lines           []InvoiceLine
	PaymentType     payment_type.PaymentType
	Note            string
}
