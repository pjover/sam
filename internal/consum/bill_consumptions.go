package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"github.com/pjover/sam/internal/domain/ports"

	"github.com/spf13/viper"
)

type BillConsumptionsManager interface {
	Run() (string, error)
}

type BillConsumptionsManagerImpl struct {
	PostManager hobbit.HttpPostManager
	dbService   ports.DbService
}

func NewBillConsumptionsManager(httpPostManager hobbit.HttpPostManager, dbService ports.DbService) BillConsumptionsManager {
	return BillConsumptionsManagerImpl{
		httpPostManager,
		dbService,
	}
}

func (b BillConsumptionsManagerImpl) Run() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.PostManager.PrettyJson(url, data)
}
