package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/storage"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type InsertConsumptionsManager struct {
	PostManager     util.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

func NewInsertConsumptionsManager() CustomerConsumptionsManager {
	return InsertConsumptionsManager{
		util.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

func (i InsertConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	child, err := i.CustomerStorage.GetChild(args.Code)
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
