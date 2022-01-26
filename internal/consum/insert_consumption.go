package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/core/ports"
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
	child, err := i.dbService.FindChild(args.Code)
	if err != nil {
		return "", err
	}
	fmt.Println("Insertant consums de l'infant", child.Name, child.Surname)

	data, err := getConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consumptions", viper.GetString("urls.hobbit"))
	return i.PostManager.PrettyJson(url, data)
}
