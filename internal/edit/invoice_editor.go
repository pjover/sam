package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/core/ports"
	"net/url"

	"github.com/spf13/viper"
)

type InvoiceEditorImpl struct {
	osService ports.OsService
}

func NewInvoiceEditor() Editor {
	return InvoiceEditorImpl{
		osService: os.NewOsService(),
	}
}

func (c InvoiceEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant la factura", code, "a", _url)
	return c.osService.OpenUrlInBrowser(_url)
}
