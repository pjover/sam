package dbo

import "github.com/pjover/sam/internal/core/model"

func ConvertProduct(product Product) model.Product {
	return model.Product{
		Id:            product.Id,
		Name:          product.Name,
		ShortName:     product.ShortName,
		Price:         product.Price,
		TaxPercentage: product.TaxPercentage,
		IsSubsidy:     product.IsSubsidy,
	}
}
