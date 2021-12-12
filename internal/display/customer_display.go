package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/spf13/viper"
)

type CustomerDisplay struct {
	getManager tuk.HttpGetManager
}

func NewCustomerDisplay() Display {
	return CustomerDisplay{
		tuk.NewHttpGetManager(),
	}
}

func (c CustomerDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/customers/%s", viper.GetString("urls.hobbit"), code)
	return c.getManager.PrettyJson(url)
}
