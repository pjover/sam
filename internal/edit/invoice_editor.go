package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"
	"net/url"

	"github.com/spf13/viper"
)

type InvoiceEditorImpl struct {
	openManager os.OpenManager
}

func NewInvoiceEditor() Editor {
	return InvoiceEditorImpl{
		openManager: os.NewOpenManager(),
	}
}

func (c InvoiceEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant la factura", code, "a", _url)
	return c.openManager.OnDefaultApp(_url)
}
