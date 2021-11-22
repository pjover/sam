package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
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
