package list

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/spf13/viper"
)

type ListCustomers struct {
	getManager tuk.HttpGetManager
}

func NewListCustomers() List {
	return ListCustomers{
		tuk.NewHttpGetManager(),
	}
}

func (l ListCustomers) List() (string, error) {
	url := fmt.Sprintf("%s/lists/customers", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
