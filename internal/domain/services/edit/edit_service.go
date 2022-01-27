package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
	"strconv"
)

type editService struct {
	externalEditor ports.ExternalEditor
}

func NewEditService(externalEditor ports.ExternalEditor) ports.EditService {
	return editService{
		externalEditor: externalEditor,
	}
}

func (e editService) EditCustomer(code int) (string, error) {
	codeStr := strconv.Itoa(code)
	url, err := e.externalEditor.Edit(ports.Customer, codeStr)
	msg := fmt.Sprintf("Editant el client %s a %s", codeStr, url)
	return msg, err
}

func (e editService) EditInvoice(code string) (string, error) {
	url, err := e.externalEditor.Edit(ports.Invoice, code)
	msg := fmt.Sprintf("Editant la factura %s a %s", code, url)
	return msg, err
}

func (e editService) EditProduct(code string) (string, error) {
	url, err := e.externalEditor.Edit(ports.Product, code)
	msg := fmt.Sprintf("Editant el producte %s a %s", code, url)
	return msg, err
}
