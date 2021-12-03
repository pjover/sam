package display

import (
	"fmt"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type ProductDisplay struct {
	getManager shared.HttpGetManager
}

func NewProductDisplay() Display {
	return ProductDisplay{
		shared.NewHttpGetManager(),
	}
}

func (p ProductDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), code)
	return p.getManager.PrettyJson(url)
}
