package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/storage"
	"github.com/spf13/viper"
)

type RectifyConsumptionsManager struct {
	PostManager     tuk.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

func NewRectifyConsumptionsManager() CustomerConsumptionsManager {
	return RectifyConsumptionsManager{
		tuk.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

func (r RectifyConsumptionsManager) Run(args CustomerConsumptionsArgs) (string, error) {
	child, err := r.CustomerStorage.GetChild(args.Code)
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
