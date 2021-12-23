package dbo

import "github.com/pjover/sam/internal/core/model"

func ConvertProduct(product Product) model.Product {
	return model.Product{
		Id:            product.Id,
		Name:          product.Name,
		ShortName:     product.ShortName,
		Price:         Decimal128ToFloat64(product.Price),
		TaxPercentage: Decimal128ToFloat64(product.TaxPercentage),
		IsSubsidy:     product.IsSubsidy,
	}
}

func ConvertProducts(products []Product) []model.Product {
	var out []model.Product
	for _, product := range products {
		out = append(out, ConvertProduct(product))
	}
	return out
}
