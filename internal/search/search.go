package search

import (
	"fmt"
	"github.com/pjover/sam/internal/adapters/hobbit"
	"net/url"

	"github.com/spf13/viper"
)

type SearchManager interface {
	SearchCustomer(args []string) (string, error)
}

type SearchManagerImpl struct {
	GetManager hobbit.HttpGetManager
}

func NewSearchManager() SearchManager {
	return SearchManagerImpl{
		GetManager: hobbit.NewHttpGetManager(),
	}
}

func (s SearchManagerImpl) SearchCustomer(args []string) (string, error) {
	text := fmt.Sprint(args)
	params := url.Values{}
	params.Add("text", text)
	_url := fmt.Sprintf("%s/search/customer?%s", viper.GetString("urls.hobbit"), params.Encode())
	return s.GetManager.PrettyJson(_url)
}
