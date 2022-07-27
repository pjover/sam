package dbo

import (
	"github.com/pjover/sam/internal/domain/model"
	"log"
)

func ConvertConsumptionsToModel(consumptions []Consumption) []model.Consumption {
	var out []model.Consumption
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumptionToModel(consumption))
	}
	return out
}

func ConvertConsumptionToModel(consumption Consumption) model.Consumption {
	return model.NewConsumption(
		consumption.Id,
		consumption.ChildId,
		consumption.ProductID,
		consumption.Units,
		convertConsumptionYearMonth(consumption.YearMonth, consumption.Id),
		consumption.Note,
		consumption.IsRectification,
		consumption.InvoiceId,
	)
}

func convertConsumptionYearMonth(yearMonth string, consumptionId string) model.YearMonth {
	ym, err := model.StringToYearMonth(yearMonth)
	if err != nil {
		log.Fatalf("la dada yearMonth '%s' del consumption %s és incorrecte", yearMonth, consumptionId)
	}
	return ym
}

func ConvertConsumptionsToDbo(consumptions []model.Consumption) []interface{} {
	var out []interface{}
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumptionToDbo(consumption))
	}
	return out
}

func ConvertConsumptionToDbo(consumption model.Consumption) Consumption {
	return Consumption{
		Id:              consumption.Id(),
		ChildId:         consumption.ChildId(),
		ProductID:       consumption.ProductId(),
		Units:           consumption.Units(),
		YearMonth:       consumption.YearMonth().String(),
		Note:            consumption.Note(),
		IsRectification: consumption.IsRectification(),
		InvoiceId:       consumption.InvoiceId(),
	}
}
