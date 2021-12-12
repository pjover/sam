package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/pjover/sam/internal/storage"
	"github.com/spf13/viper"
)

type BillConsumptionsManager interface {
	Run() (string, error)
}

type BillConsumptionsManagerImpl struct {
	PostManager     tuk.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

func NewBillConsumptionsManager() BillConsumptionsManager {
	return BillConsumptionsManagerImpl{
		tuk.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

func (b BillConsumptionsManagerImpl) Run() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.PostManager.PrettyJson(url, data)
}
