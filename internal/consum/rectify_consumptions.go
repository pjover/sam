package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/viper"
)

type RectifyConsumptionsManager struct {
	PostManager tuk.HttpPostManager
	dbService   ports.DbService
}

func NewRectifyConsumptionsManager() CustomerConsumptionsManager {
	return RectifyConsumptionsManager{
		tuk.NewHttpPostManager(),
		mongo_db.NewDbService(cfg.NewConfigService()),
	}
}

func (r RectifyConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	child, err := r.dbService.GetChild(args.Code)
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
