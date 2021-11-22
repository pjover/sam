package list

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type ListCustomers struct {
	getManager util.HttpGetManager
}

func NewListCustomers() List {
	return ListCustomers{
		util.NewHttpGetManager(),
	}
}

func (l ListCustomers) List() (string, error) {
	url := fmt.Sprintf("%s/lists/customers", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
