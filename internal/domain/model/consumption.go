package model

import (
	"bytes"
	"fmt"
)

const ConsumptionIdLength = 8

type Consumption struct {
	id              string
	childId         int
	productId       string
	units           float64
	yearMonth       YearMonth
	note            string
	isRectification bool
	invoiceId       string
}

func NewConsumption(
	id string,
	childId int,
	productId string,
	units float64,
	yearMonth YearMonth,
	note string,
	isRectification bool,
	invoiceId string,
) Consumption {
	return Consumption{
		id:              id,
		childId:         childId,
		productId:       productId,
		units:           units,
		yearMonth:       yearMonth,
		note:            note,
		isRectification: isRectification,
		invoiceId:       invoiceId,
	}
}

func (c Consumption) Id() string {
	return c.id
}

func (c Consumption) ChildId() int {
	return c.childId
}

func (c Consumption) ProductId() string {
	return c.productId
}

func (c Consumption) Units() float64 {
	return c.units
}

func (c Consumption) YearMonth() YearMonth {
	return c.yearMonth
}

func (c Consumption) Note() string {
	return c.note
}

func (c Consumption) IsRectification() bool {
	return c.isRectification
}

func (c Consumption) InvoiceId() string {
	return c.invoiceId
}

func (c Consumption) String() string {
	return fmt.Sprintf("%d  %s  %4.1f  %s  %-5v  %s  %s", c.ChildId(), c.YearMonth(), c.Units(), c.ProductId(), c.IsRectification(), c.InvoiceId(), c.Note())
}

func (c Consumption) CopyWithNewInvoiceId(invoiceId string) Consumption {
	return Consumption{
		id:              c.id,
		childId:         c.childId,
		productId:       c.productId,
		units:           c.units,
		yearMonth:       c.yearMonth,
		note:            c.note,
		isRectification: c.isRectification,
		invoiceId:       invoiceId,
	}
}

func ConsumptionListToString(consumptions []Consumption, child Child, products map[string]Product) string {
	var total float64
	var buffer bytes.Buffer
	for _, c := range consumptions {
		if c.ChildId() != child.Id() {
			continue
		}
		product := products[c.ProductId()]
		price := c.Units() * product.Price()
		total += price
		buffer.WriteString(fmt.Sprintf("  [%s]  %4.1f x %s : %7.2f\n",
			c.YearMonth().String(),
			c.Units(),
			c.ProductId(),
			price,
		))
	}
	title := fmt.Sprintf("%s: %.02f €\n", child.NameWithId(), total)
	return title + buffer.String()
}
