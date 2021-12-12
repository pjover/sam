package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/viper"
)

type InsertConsumptionsManager struct {
	PostManager tuk.HttpPostManager
	dbService   ports.DbService
}

func NewInsertConsumptionsManager() CustomerConsumptionsManager {
	return InsertConsumptionsManager{
		tuk.NewHttpPostManager(),
		mongo_db.NewDbService(),
	}
}

func (i InsertConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	child, err := i.dbService.GetChild(args.Code)
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
