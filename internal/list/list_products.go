package list

import (
	"fmt"

	"github.com/pjover/sam/internal/shared"
	"github.com/spf13/viper"
)

type ListProducts struct {
	getManager shared.HttpGetManager
}

func NewListProducts() List {
	return ListProducts{
		shared.NewHttpGetManager(),
	}
}

func (l ListProducts) List() (string, error) {
	fmt.Println("Llistat de tots els productes")
	url := fmt.Sprintf("%s/products?page=0&size=100", viper.GetString("urls.hobbit"))
	return l.getManager.PrettyJson(url)
}
