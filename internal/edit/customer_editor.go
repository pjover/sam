package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"

	"github.com/spf13/viper"
)

type CustomerEditorImpl struct {
	openManager os.OpenManager
}

func NewCustomerEditor() Editor {
	return CustomerEditorImpl{
		openManager: os.NewOpenManager(),
	}
}

func (c CustomerEditorImpl) Edit(code string) error {
	_url := fmt.Sprintf("%s/customer/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el client", code, "a", _url)
	return c.openManager.OnDefaultApp(_url)
}
