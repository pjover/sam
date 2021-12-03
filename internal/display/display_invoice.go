package display

import (
	"fmt"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type InvoiceDisplay struct {
	getManager shared.HttpGetManager
}

func NewInvoiceDisplay() Display {
	return InvoiceDisplay{
		shared.NewHttpGetManager(),
	}
}

func (i InvoiceDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), code)
	return i.getManager.PrettyJson(url)
}
