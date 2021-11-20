package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type ListChildren interface {
	List() (string, error)
}

type ListChildrenImpl struct {
	getManager util.HttpGetManager
}

func NewListChildren() ListChildren {
	return ListChildrenImpl{
		util.NewHttpGetManager(),
	}
}

func (l ListChildrenImpl) List() (string, error) {
	url := fmt.Sprintf("%s/lists/children", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
