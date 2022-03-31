package model

import "fmt"

type InvoiceLine struct {
	ProductId     string
	Units         float64
	ProductPrice  float64
	TaxPercentage float64
	ChildId       int
}

func (i InvoiceLine) Format() string {
	return fmt.Sprintf(
		"%.1f %s (%.2f)",
		i.Units,
		i.ProductId,
		i.Units*i.ProductPrice,
	)
}
