package dbo

import "github.com/pjover/sam/internal/domain/model"

func ConvertConsumptionToModel(consumption Consumption) model.Consumption {
	return model.Consumption{
		Id:              consumption.Id,
		ChildId:         consumption.ChildId,
		ProductId:       consumption.ProductID,
		Units:           Decimal128ToFloat64(consumption.Units),
		YearMonth:       consumption.YearMonth,
		Note:            consumption.Note,
		IsRectification: consumption.IsRectification,
		InvoiceId:       consumption.InvoiceId,
	}
}

func ConvertConsumptionsToModel(consumptions []Consumption) []model.Consumption {
	var out []model.Consumption
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumptionToModel(consumption))
	}
	return out
}

func ConvertConsumptionToDbo(consumption model.Consumption) Consumption {
	return Consumption{
		Id:              consumption.Id,
		ChildId:         consumption.ChildId,
		ProductID:       consumption.ProductId,
		Units:           Float64ToDecimal128(consumption.Units),
		YearMonth:       consumption.YearMonth,
		Note:            consumption.Note,
		IsRectification: consumption.IsRectification,
		InvoiceId:       consumption.InvoiceId,
	}
}

func ConvertConsumptionsToDbo(consumptions []model.Consumption) []interface{} {
	var out []interface{}
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumptionToDbo(consumption))
	}
	return out
}
