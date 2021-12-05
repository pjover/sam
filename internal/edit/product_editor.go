package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"
	"net/url"

	"github.com/spf13/viper"
)

type ProductEditorImpl struct {
	openManager os.OpenManager
}

func NewProductEditor() Editor {
	return ProductEditorImpl{
		openManager: os.NewOpenManager(),
	}
}

func (c ProductEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/product/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant el producte", code, "a", _url)
	return c.openManager.OnDefaultApp(_url)
}
