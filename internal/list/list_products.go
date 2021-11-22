package list

import (
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
)

type ListProducts struct {
	getManager util.HttpGetManager
}

func NewListProducts() List {
	return ListProducts{
		util.NewHttpGetManager(),
	}
}

func (l ListProducts) List() (string, error) {
	fmt.Println("Llistat de tots els productes")
	url := fmt.Sprintf("%s/products?page=0&size=100", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
