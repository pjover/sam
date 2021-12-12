package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/core/ports"

	"github.com/spf13/viper"
)

type CustomerEditorImpl struct {
	osService ports.OsService
}

func NewCustomerEditor() Editor {
	return CustomerEditorImpl{
		osService: os.NewOsService(),
	}
}

func (c CustomerEditorImpl) Edit(code string) error {
	_url := fmt.Sprintf("%s/customer/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el client", code, "a", _url)
	return c.osService.OpenUrlInBrowser(_url)
}
