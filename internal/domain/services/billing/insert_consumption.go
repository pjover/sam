package billing

import (
	"bytes"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"github.com/spf13/viper"
)

type BillingService interface {
	InsertConsumptions(Code int, Consumptions map[string]float64, Note string) (string, error)
}

type billingService struct {
	dbService ports.DbService
}

func NewBillingService(dbService ports.DbService) BillingService {
	return billingService{
		dbService,
	}
}

func (i billingService) InsertConsumptions(code int, consumptions map[string]float64, note string) (string, error) {
	var buffer bytes.Buffer

	child, err := i.dbService.FindChild(code)
	if err != nil {
		return "", err
	}
	products, err := i.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	yearMonth := viper.GetString("yearMonth")
	var first = true
	var completeConsumptions []model.Consumption
	for id, units := range consumptions {
		c := model.Consumption{
			Code:      common.RandString(model.ConsumptionCodeLength),
			ChildCode: code,
			ProductID: id,
			Units:     units,
			YearMonth: yearMonth,
		}
		if first {
			c.Note = note
			first = false
		}
		completeConsumptions = append(completeConsumptions, c)
	}

	err = i.dbService.InsertConsumptions(completeConsumptions)
	if err != nil {
		return "", err
	}

	buffer.WriteString(model.ConsumptionListToString(completeConsumptions, child, products))

	return buffer.String(), nil
}
