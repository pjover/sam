package adm

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/util"
)

type SearchManager struct {
	GetManager util.HttpGetManager
}

func NewSearchManager() SearchManager {
	return SearchManager{
		GetManager: new(util.SamHttpGetManager),
	}
}

func (s SearchManager) SearchCustomer(args []string) (string, error) {
	text := fmt.Sprint(args)
	params := url.Values{}
	params.Add("text", text)
	_url := fmt.Sprintf("%s/search/customer?%s", viper.GetString("urls.hobbit"), params.Encode())
	return s.GetManager.GetPrint(_url)
}
