package edit

import (
	"fmt"
	"net/url"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type ProductEditorImpl struct {
}

func NewProductEditor() Editor {
	return ProductEditorImpl{}
}

func (c ProductEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/product/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant el producte", code, "a", _url)
	return shared.OpenOnDefaultApp(_url)
}
