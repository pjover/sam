package list

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type ListChildren struct {
	getManager util.HttpGetManager
}

func NewListChildren() List {
	return ListChildren{
		util.NewHttpGetManager(),
	}
}

func (l ListChildren) List() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
