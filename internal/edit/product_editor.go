package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/core/ports"
	"net/url"

	"github.com/spf13/viper"
)

type ProductEditorImpl struct {
	osService ports.OsService
}

func NewProductEditor() Editor {
	return ProductEditorImpl{
		osService: os.NewOsService(),
	}
}

func (c ProductEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/product/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant el producte", code, "a", _url)
	return c.osService.OpenUrlInBrowser(_url)
}
