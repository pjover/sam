package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/viper"
)

type RectifyConsumptionsManager struct {
	PostManager hobbit.HttpPostManager
	dbService   ports.DbService
}

func NewRectifyConsumptionsManager(httpPostManager hobbit.HttpPostManager, dbService ports.DbService) CustomerConsumptionsManager {
	return RectifyConsumptionsManager{
		httpPostManager,
		dbService,
	}
}

func (r RectifyConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	child, err := r.dbService.FindChild(args.Id)
	if err != nil {
		return "", err
	}
	fmt.Println("Rectificant els consums de l'infant", child.Name, child.Surname)

	data, err := getConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consumptions/rectification", viper.GetString("urls.hobbit"))
	return r.PostManager.PrettyJson(url, data)
}
