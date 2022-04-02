package model

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	id            string
	name          string
	shortName     string
	price         float64
	taxPercentage float64
	isSubsidy     bool
}

func (p Product) Id() string {
	return p.id
}

func (p Product) Name() string {
	return p.name
}

func (p Product) ShortName() string {
	return p.shortName
}

func (p Product) Price() float64 {
	return p.price
}

func (p Product) TaxPercentage() float64 {
	return p.taxPercentage
}

func (p Product) IsSubsidy() bool {
	return p.isSubsidy
}

func (p Product) String() string {
	return fmt.Sprintf("%s  % 7.2f  %s", p.id, p.price, p.name)
}

func (p Product) validate() error {
	if p.id == "" {
		return errors.New("el id del producte no pot estar buit")
	}

	if p.id != strings.ToUpper(p.id) {
		return errors.New("el id del producte ha d'estar en majúscules")
	}

	if len(p.id) != 3 {
		return errors.New("el id del producte ha de tenir 3 caràcters")
	}

	if p.shortName == "" {
		return errors.New("el nom curt del producte (ShortName) no pot estar buit")
	}

	if p.name == "" {
		return errors.New("el nom del producte (Name) no pot estar buit")
	}

	if p.price == 0 {
		return errors.New("el preu del producte (Price) no pot ser 0.0")
	}
	return nil
}

func NewProduct(id string, name string, shortName string, price float64, taxPercentage float64, isSubsidy bool) (Product, error) {
	product := Product{
		id:            id,
		name:          name,
		shortName:     shortName,
		price:         price,
		taxPercentage: taxPercentage,
		isSubsidy:     isSubsidy,
	}
	return product, product.validate()
}

func GetProduct(productID string, products []Product) Product {
	for _, product := range products {
		if product.Id() == productID {
			return product
		}
	}
	return Product{}
}
