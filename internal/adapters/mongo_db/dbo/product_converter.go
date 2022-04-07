package dbo

import "github.com/pjover/sam/internal/domain/model"

func ConvertProductToModel(product Product) (model.Product, error) {
	return model.NewProduct(
		product.Id,
		product.Name,
		product.ShortName,
		Decimal128ToFloat64(product.Price),
		Decimal128ToFloat64(product.TaxPercentage),
		product.IsSubsidy,
	)
}

func ConvertProductsToModel(products []Product) ([]model.Product, error) {
	var out []model.Product
	for _, product := range products {
		prod, err := ConvertProductToModel(product)
		if err != nil {
			return nil, err
		}
		out = append(out, prod)
	}
	return out, nil
}

func ConvertProductToDbo(product model.Product) Product {
	return Product{
		Id:            product.Id(),
		Name:          product.Name(),
		ShortName:     product.ShortName(),
		Price:         Float64ToDecimal128(product.Price()),
		TaxPercentage: Float64ToDecimal128(product.TaxPercentage()),
		IsSubsidy:     product.IsSubsidy(),
	}
}
