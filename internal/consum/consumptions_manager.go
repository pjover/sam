package consum

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sam/internal/util"
	"sam/storage"
)

type ConsumptionsManager interface {
	InsertConsumptions(args InsertConsumptionsArgs) (string, error)
	RectifyConsumptions(args InsertConsumptionsArgs) (string, error)
	BillConsumptions() (string, error)
}

type ConsumptionsManagerImpl struct {
	PostManager     util.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

type InsertConsumptionsArgs struct {
	Code         int
	Consumptions map[string]float64
	Note         string
}

func NewConsumptionsManager() ConsumptionsManager {
	return ConsumptionsManagerImpl{
		util.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

func (c ConsumptionsManagerImpl) InsertConsumptions(args InsertConsumptionsArgs) (string, error) {
	child, err := c.CustomerStorage.GetChild(args.Code)
	if err != nil {
		return "", err
	}
	fmt.Println("Insertant consums de l'infant", child.Name, child.Surname)

	data, err := c.getInsertConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consum", viper.GetString("urls.hobbit"))
	return c.PostManager.PrettyJson(url, data)
}

type ConsumptionBody struct {
	ProductID string  `json:"productId"`
	Units     float64 `json:"units"`
	Note      string  `json:"note"`
}
type ChildBody struct {
	Code         int               `json:"code"`
	Consumptions []ConsumptionBody `json:"consum"`
}
type Body struct {
	YearMonth string      `json:"yearMonth"`
	Children  []ChildBody `json:"children"`
}

func (c ConsumptionsManagerImpl) RectifyConsumptions(args InsertConsumptionsArgs) (string, error) {
	child, err := c.CustomerStorage.GetChild(args.Code)
	if err != nil {
		return "", err
	}
	fmt.Println("Rectificant els consums de l'infant", child.Name, child.Surname)

	data, err := c.getInsertConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consum/rectification", viper.GetString("urls.hobbit"))
	return c.PostManager.PrettyJson(url, data)
}

func (c ConsumptionsManagerImpl) getInsertConsumptionsJson(args InsertConsumptionsArgs) ([]byte, error) {
	consumptionsBody := c.buildConsumptionsBody(args)
	bytes, err := json.Marshal(consumptionsBody)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}

func (c ConsumptionsManagerImpl) buildConsumptionsBody(args InsertConsumptionsArgs) Body {
	var consumptions []ConsumptionBody
	var note = args.Note
	for key, val := range args.Consumptions {
		consumptions = append(consumptions, ConsumptionBody{key, val, note})
		note = ""
	}
	consumptionsBody := Body{
		YearMonth: viper.GetString("yearMonth"),
		Children: []ChildBody{
			{args.Code, consumptions},
		},
	}
	return consumptionsBody
}

func (c ConsumptionsManagerImpl) BillConsumptions() (string, error) {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")

	url := fmt.Sprintf("%s/billing/billConsumptions", viper.GetString("urls.hobbit"))
	var data []byte
	return c.PostManager.PrettyJson(url, data)
}
