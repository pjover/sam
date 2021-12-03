package list

import (
	"fmt"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type ListChildren struct {
	getManager shared.HttpGetManager
}

func NewListChildren() List {
	return ListChildren{
		shared.NewHttpGetManager(),
	}
}

func (l ListChildren) List() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
