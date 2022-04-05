package model

import (
	"bytes"
	"fmt"
)

const ConsumptionIdLength = 8

type Consumption struct {
	Id              string
	ChildId         int
	ProductId       string
	Units           float64
	YearMonth       YearMonth
	Note            string
	IsRectification bool
	InvoiceId       string
}

func (c Consumption) String() string {
	return fmt.Sprintf("%d  %s  %4.1f  %s  %-5v  %s  %s", c.ChildId, c.YearMonth, c.Units, c.ProductId, c.IsRectification, c.InvoiceId, c.Note)
}

func ConsumptionListToString(consumptions []Consumption, child Child, products map[string]Product) string {
	var total float64
	var buffer bytes.Buffer
	for _, c := range consumptions {
		if c.ChildId != child.Id() {
			continue
		}
		product := products[c.ProductId]
		price := c.Units * product.Price()
		total += price
		buffer.WriteString(fmt.Sprintf("  [%s]  %4.1f x %s : %7.2f\n",
			c.YearMonth.String(),
			c.Units,
			c.ProductId,
			price,
		))
	}
	title := fmt.Sprintf("%s: %.02f €\n", child.NameWithId(), total)
	return title + buffer.String()
}
