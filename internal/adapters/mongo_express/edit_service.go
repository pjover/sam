package mongo_express

import (
	"fmt"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/spf13/viper"
	"net/url"
)

type editService struct {
	osService ports.OsService
}

func NewEditService(osService ports.OsService) ports.EditService {
	return editService{
		osService: osService,
	}
}

func (e editService) EditCustomer(code int) (string, error) {
	editUrl := fmt.Sprintf("%s/customer/%d", viper.GetString("urls.mongoExpress"), code)
	msg := fmt.Sprintf("Editant el client %d a %s", code, editUrl)
	return msg, e.osService.OpenUrlInBrowser(editUrl)
}

func (e editService) EditInvoice(code string) (string, error) {
	escapedCode := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	editUrl := fmt.Sprintf("%s/invoice/%s", viper.GetString("urls.mongoExpress"), escapedCode)
	msg := fmt.Sprintf("Editant la factura %s a %s", code, editUrl)
	return msg, e.osService.OpenUrlInBrowser(editUrl)
}

func (e editService) EditProduct(code string) (string, error) {
	escapedCode := url.QueryEscape(fmt.Sprintf("\"%s\"", code))
	editUrl := fmt.Sprintf("%s/product/%s", viper.GetString("urls.mongoExpress"), escapedCode)
	msg := fmt.Sprintf("Editant el producte %s a %s", code, editUrl)
	return msg, e.osService.OpenUrlInBrowser(editUrl)
}
