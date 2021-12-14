package consum

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/mongo_db"
	"github.com/pjover/sam/internal/adapters/tuk"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/viper"
)

type BillConsumptionsManager interface {
	Run() (string, error)
}

type BillConsumptionsManagerImpl struct {
	PostManager tuk.HttpPostManager
	dbService   ports.DbService
}

func NewBillConsumptionsManager() BillConsumptionsManager {
	return BillConsumptionsManagerImpl{
		tuk.NewHttpPostManager(),
		mongo_db.NewDbService(),
	}
}

func (b BillConsumptionsManagerImpl) Run() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return b.PostManager.PrettyJson(url, data)
}
