package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type ListMails interface {
	List(ei1 bool, ei2 bool, ei3 bool) (string, error)
}

type ListMailsImpl struct {
	getManager util.HttpGetManager
}

func NewListMails() ListMails {
	return ListMailsImpl{
		util.NewHttpGetManager(),
	}
}

func (l ListMailsImpl) List(ei1 bool, ei2 bool, ei3 bool) (string, error) {
	var group string
	if ei1 {
		group = "EI_1"
	} else if ei2 {
		group = "EI_2"
	} else if ei3 {
		group = "EI_3"
	} else {
		group = "ALL"
	}

	url := fmt.Sprintf("%s/lists/emails/%s", viper.GetString("urls.hobbit"), group)
	return l.getManager.PrettyJson(url)
}
