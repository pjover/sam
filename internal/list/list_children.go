package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"

	"github.com/spf13/viper"
)

type ListChildren struct {
	getManager hobbit.HttpGetManager
}

func NewListChildren() List {
	return ListChildren{
		hobbit.NewHttpGetManager(),
	}
}

func (l ListChildren) List() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
