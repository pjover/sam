package model

import "fmt"

type Product struct {
	Id            string
	Name          string
	ShortName     string
	Price         float64
	TaxPercentage float64
	IsSubsidy     bool
}

func (p Product) String() string {
	return fmt.Sprintf("%s  % 7.2f  %s", p.Id, p.Price, p.Name)
}

func GetProduct(productID string, products []Product) Product {
	for _, product := range products {
		if product.Id == productID {
			return product
		}
	}
	return Product{}
}
