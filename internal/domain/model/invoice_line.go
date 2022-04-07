package model

import "fmt"

type InvoiceLine struct {
	productId     string
	units         float64
	productPrice  float64
	taxPercentage float64
	childId       int
}

func NewInvoiceLine(
	productId string,
	units float64,
	productPrice float64,
	taxPercentage float64,
	childId int,
) InvoiceLine {
	return InvoiceLine{
		productId:     productId,
		units:         units,
		productPrice:  productPrice,
		taxPercentage: taxPercentage,
		childId:       childId,
	}
}

func (i InvoiceLine) ProductId() string {
	return i.productId
}

func (i InvoiceLine) Units() float64 {
	return i.units
}

func (i InvoiceLine) ProductPrice() float64 {
	return i.productPrice
}

func (i InvoiceLine) TaxPercentage() float64 {
	return i.taxPercentage
}

func (i InvoiceLine) ChildId() int {
	return i.childId
}

func (i InvoiceLine) Format() string {
	return fmt.Sprintf(
		"%.1f %s (%.2f)",
		i.units,
		i.productId,
		i.units*i.productPrice,
	)
}
