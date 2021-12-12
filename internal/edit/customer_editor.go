package edit

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"

	"github.com/spf13/viper"
)

type CustomerEditorImpl struct {
	execManager os.ExecManager
}

func NewCustomerEditor() Editor {
	return CustomerEditorImpl{
		execManager: os.NewExecManager(),
	}
}

func (c CustomerEditorImpl) Edit(code string) error {
	_url := fmt.Sprintf("%s/customer/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el client", code, "a", _url)
	return c.execManager.BrowseTo(_url)
}
