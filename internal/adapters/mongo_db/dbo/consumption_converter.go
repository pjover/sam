package dbo

import "github.com/pjover/sam/internal/core/model"

func ConvertConsumption(consumption Consumption) model.Consumption {
	return model.Consumption{
		Code:            consumption.Code,
		ChildCode:       consumption.ChildCode,
		ProductID:       consumption.ProductID,
		Units:           Decimal128ToFloat64(consumption.Units),
		YearMonth:       consumption.YearMonth,
		IsRectification: consumption.IsRectification,
		InvoiceCode:     consumption.InvoiceCode,
	}
}

func ConvertConsumptions(consumptions []Consumption) []model.Consumption {
	var out []model.Consumption
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumption(consumption))
	}
	return out
}
