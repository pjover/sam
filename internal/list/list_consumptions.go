package list

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type ListConsumptions interface {
	ListOne(customerCode int) (string, error)
	List() (string, error)
}

type ListConsumptionsImpl struct {
	getManager util.HttpGetManager
}

func NewListConsumptions() ListConsumptions {
	return ListConsumptionsImpl{
		util.NewHttpGetManager(),
	}
}

func (l ListConsumptionsImpl) ListOne(customerCode int) (string, error) {
	fmt.Println("Llistat dels consums pendents del client", customerCode)
	url := fmt.Sprintf("%s/consumptions/%d", viper.GetString("urls.hobbit"), customerCode)
	return l.getManager.PrettyJson(url)
}

func (l ListConsumptionsImpl) List() (string, error) {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	url := fmt.Sprintf("%s/consumptions", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
