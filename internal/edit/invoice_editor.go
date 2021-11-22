package edit

import (
	"fmt"
	"net/url"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type InvoiceEditorImpl struct {
}

func NewInvoiceEditor() Editor {
	return InvoiceEditorImpl{}
}

func (c InvoiceEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant la factura", code, "a", _url)
	return util.OpenOnBrowser(_url)
}
