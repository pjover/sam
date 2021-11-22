package search

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
	"sam/internal/util"
)

type SearchManager interface {
	SearchCustomer(args []string) (string, error)
}

type SearchManagerImpl struct {
	GetManager util.HttpGetManager
}

func NewSearchManager() SearchManager {
	return SearchManagerImpl{
		GetManager: util.NewHttpGetManager(),
	}
}

func (s SearchManagerImpl) SearchCustomer(args []string) (string, error) {
	text := fmt.Sprint(args)
	params := url.Values{}
	params.Add("text", text)
	_url := fmt.Sprintf("%s/search/customer?%s", viper.GetString("urls.hobbit"), params.Encode())
	return s.GetManager.PrettyJson(_url)
}
