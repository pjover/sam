package consum

import (
	"bytes"
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/common"
	"github.com/spf13/viper"
)

type InsertConsumptionsManager struct {
	PostManager hobbit.HttpPostManager
	dbService   ports.DbService
}

func NewInsertConsumptionsManager(httpPostManager hobbit.HttpPostManager, dbService ports.DbService) CustomerConsumptionsManager {
	return InsertConsumptionsManager{
		httpPostManager,
		dbService,
	}
}

func (i InsertConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	var buffer bytes.Buffer

	child, err := i.dbService.FindChild(args.Code)
	if err != nil {
		return "", err
	}
	buffer.WriteString(fmt.Sprintf("Insertant consums de l'infant %s\n", child))

	yearMonth := viper.GetString("yearMonth")
	var first = true
	var consumptions []model.Consumption
	for id, units := range args.Consumptions {
		c := model.Consumption{
			Code:      common.RandString(model.ConsumptionCodeLength),
			ChildCode: args.Code,
			ProductID: id,
			Units:     units,
			YearMonth: yearMonth,
		}
		if first {
			c.Note = args.Note
			first = false
		}
		consumptions = append(consumptions, c)
	}

	err = i.dbService.InsertConsumptions(consumptions)
	if err != nil {
		return "", err
	}

	for _, consumption := range consumptions {
		buffer.WriteString(consumption.String() + "\n")
	}

	return buffer.String(), nil
}
