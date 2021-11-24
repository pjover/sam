package consum

import (
	"fmt"

	"github.com/pjover/sam/internal/storage"
	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type ConsumptionsManager interface {
	BillConsumptions() (string, error)
}

type ConsumptionsManagerImpl struct {
	PostManager     util.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

func NewConsumptionsManager() ConsumptionsManager {
	return ConsumptionsManagerImpl{
		util.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

func (c ConsumptionsManagerImpl) BillConsumptions() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return c.PostManager.PrettyJson(url, data)
}
