package model

import (
	"fmt"
)

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
