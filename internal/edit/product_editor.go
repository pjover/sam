package edit

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/internal/util"
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
	return util.OpenOnBrowser(_url)
}
