package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/spf13/viper"
)

type ListConsumptions interface {
	ListOne(childCode int) (string, error)
	List() (string, error)
}

type ListConsumptionsImpl struct {
	getManager tuk.HttpGetManager
}

func NewListConsumptions() ListConsumptions {
	return ListConsumptionsImpl{
		tuk.NewHttpGetManager(),
	}
}

func (l ListConsumptionsImpl) ListOne(childCode int) (string, error) {
	fmt.Println("Llistat dels consums pendents del client", childCode)
	url := fmt.Sprintf("%s/consumptions/%d", viper.GetString("urls.hobbit"), childCode)
	return l.getManager.PrettyJson(url)
}

func (l ListConsumptionsImpl) List() (string, error) {
	fmt.Println("Llistat dels consums pendents de tots els clients")
	url := fmt.Sprintf("%s/consumptions", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
