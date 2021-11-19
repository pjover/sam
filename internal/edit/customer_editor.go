package edit

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type CustomerEditorImpl struct {
}

func newCustomerEditor() Editor {
	return CustomerEditorImpl{}
}

func (c CustomerEditorImpl) Edit(code string) error {
	_url := fmt.Sprintf("%s/customer/%s", viper.GetString("urls.mongoExpress"), code)
	fmt.Println("Editant el client", code, "a", _url)
	return util.OpenOnBrowser(_url)
}
