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

func (e editService) EditCustomer(id int) (string, error) {
	idStr := strconv.Itoa(id)
	url, err := e.externalEditor.Edit(ports.Customer, idStr)
	msg := fmt.Sprintf("Editant el client %s a %s", idStr, url)
	return msg, err
}

func (e editService) EditInvoice(id string) (string, error) {
	url, err := e.externalEditor.Edit(ports.Invoice, id)
	msg := fmt.Sprintf("Editant la factura %s a %s", id, url)
	return msg, err
}

func (e editService) EditProduct(id string) (string, error) {
	url, err := e.externalEditor.Edit(ports.Product, id)
	msg := fmt.Sprintf("Editant el producte %s a %s", id, url)
	return msg, err
}
