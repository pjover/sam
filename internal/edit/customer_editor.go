package edit

import (
	"fmt"

	"github.com/pjover/sam/internal/util"
	"github.com/spf13/viper"
)

type CustomerEditorImpl struct {
}

func NewCustomerEditor() Editor {
	return CustomerEditorImpl{}
}

func (c CustomerEditorImpl) Edit(code string) error {
	_url := fmt.Sprintf("%s/customer/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el client", code, "a", _url)
	return util.OpenOnBrowser(_url)
}
