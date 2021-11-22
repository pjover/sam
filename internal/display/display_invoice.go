package display

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type InvoiceDisplay struct {
	getManager util.HttpGetManager
}

func NewInvoiceDisplay() Display {
	return InvoiceDisplay{
		util.NewHttpGetManager(),
	}
}

func (i InvoiceDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), code)
	return i.getManager.PrettyJson(url)
}
