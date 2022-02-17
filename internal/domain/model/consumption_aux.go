package model

import (
	"bytes"
	"fmt"
)

func (c Consumption) String() string {
	return fmt.Sprintf("%d  %s  %4.1f  %s  %-5v  %s  %s", c.ChildCode, c.YearMonth, c.Units, c.ProductId, c.IsRectification, c.InvoiceCode, c.Note)
}

func ConsumptionListToString(consumptions []Consumption, child Child, products []Product) string {
	var total float64
	var buffer bytes.Buffer
	for _, c := range consumptions {
		if c.ChildCode != child.Id {
			continue
		}
		product := GetProduct(c.ProductId, products)
		price := c.Units * product.Price
		total += price
		buffer.WriteString(fmt.Sprintf("  [%s]  %4.1f x %s : % 7.2f\n",
			c.YearMonth,
			c.Units,
			c.ProductId,
			price,
		))
	}
	title := fmt.Sprintf("%s: %.02f €\n", child.NameWithCode(), total)
	return title + buffer.String()
}
