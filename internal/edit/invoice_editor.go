package edit

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/internal/util"
)

type InvoiceEditorImpl struct {
}

func newInvoiceEditor() Editor {
	return InvoiceEditorImpl{}
}

func (c InvoiceEditorImpl) Edit(code string) error {
	_code := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	_url := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), _code)
	fmt.Println("Editant la factura", code, "a", _url)
	return util.OpenOnBrowser(_url)
}
