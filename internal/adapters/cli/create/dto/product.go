package dto

import "github.com/pjover/sam/internal/domain/model"

type Product struct {
	Id            string
	Name          string
	ShortName     string
	Price         float64
	TaxPercentage float64
	IsSubsidy     bool
}

func ProductToModel(product Product) (model.Product, error) {
	return model.NewProduct(
		product.Id,
		product.Name,
		product.ShortName,
		product.Price,
		product.TaxPercentage,
		product.IsSubsidy,
	)
}
