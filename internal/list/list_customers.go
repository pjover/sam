package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"

	"github.com/spf13/viper"
)

type ListCustomers struct {
	getManager hobbit.HttpGetManager
}

func NewListCustomers() List {
	return ListCustomers{
		hobbit.NewHttpGetManager(),
	}
}

func (l ListCustomers) List() (string, error) {
	url := fmt.Sprintf("%s/lists/customers", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
