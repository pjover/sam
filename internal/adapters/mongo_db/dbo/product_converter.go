package dbo

import "github.com/pjover/sam/internal/domain/model"

func ConvertProductToModel(product Product) model.Product {
	return model.Product{
		Id:            product.Id,
		Name:          product.Name,
		ShortName:     product.ShortName,
		Price:         Decimal128ToFloat64(product.Price),
		TaxPercentage: Decimal128ToFloat64(product.TaxPercentage),
		IsSubsidy:     product.IsSubsidy,
	}
}

func ConvertProductsToModel(products []Product) []model.Product {
	var out []model.Product
	for _, product := range products {
		out = append(out, ConvertProductToModel(product))
	}
	return out
}

func ConvertProductToDbo(product model.Product) Product {
	return Product{
		Id:            product.Id,
		Name:          product.Name,
		ShortName:     product.ShortName,
		Price:         Float64ToDecimal128(product.Price),
		TaxPercentage: Float64ToDecimal128(product.TaxPercentage),
		IsSubsidy:     product.IsSubsidy,
	}
}
