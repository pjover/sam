package display

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type ProductDisplay struct {
	getManager util.HttpGetManager
}

func NewProductDisplay() Display {
	return ProductDisplay{
		util.NewHttpGetManager(),
	}
}

func (p ProductDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/products/%s", viper.GetString("urls.hobbit"), code)
	return p.getManager.PrettyJson(url)
}
