package display

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/tuk"

	"github.com/spf13/viper"
)

type InvoiceDisplay struct {
	getManager tuk.HttpGetManager
}

func NewInvoiceDisplay() Display {
	return InvoiceDisplay{
		tuk.NewHttpGetManager(),
	}
}

func (i InvoiceDisplay) Display(code string) (string, error) {
	url := fmt.Sprintf("%s/invoices/%s", viper.GetString("urls.hobbit"), code)
	return i.getManager.PrettyJson(url)
}
