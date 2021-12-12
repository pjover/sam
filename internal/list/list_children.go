package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/spf13/viper"
)

type ListChildren struct {
	getManager tuk.HttpGetManager
}

func NewListChildren() List {
	return ListChildren{
		tuk.NewHttpGetManager(),
	}
}

func (l ListChildren) List() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
