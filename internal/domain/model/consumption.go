package model

import (
	"bytes"
	"fmt"
	"sort"
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

func ConsumptionListFormatValues(consumptions []Consumption, child Child, products map[string]Product, indentText string) (string, float64) {
	var total float64
	var buffer bytes.Buffer

	sort.Slice(consumptions, func(i, j int) bool {
		if consumptions[i].childId != consumptions[j].childId {
			return consumptions[i].childId < consumptions[j].childId
		} else {
			return consumptions[i].productId < consumptions[j].productId
		}
	})

	for _, c := range consumptions {
		if c.ChildId() != child.Id() {
			continue
		}
		product := products[c.ProductId()]
		price := c.Units() * product.Price()
		total += price
		buffer.WriteString(fmt.Sprintf("%s  [%s]  %4.1f x %s : %7.2f\n",
			indentText,
			c.YearMonth().String(),
			c.Units(),
			c.ProductId(),
			price,
		))
	}
	title := fmt.Sprintf("%s%s: %.02f €\n", indentText, child.NameWithId(), total)
	return title + buffer.String(), total
}

type TransientConsumption struct {
	Id              string
	ChildId         int
	ProductId       string
	Units           float64
	YearMonth       YearMonth
	Note            string
	IsRectification bool
}
