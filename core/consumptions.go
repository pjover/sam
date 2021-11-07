package core

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sam/storage"
	"sam/util"
)

type ConsumptionsManager struct {
	PostManager     util.HttpPostManager
	CustomerStorage storage.CustomerStorage
}

func NewConsumptionsManager() ConsumptionsManager {
	return ConsumptionsManager{
		util.NewHttpPostManager(),
		storage.NewCustomerStorage(),
	}
}

type InsertConsumptionsArgs struct {
	Code         int
	Consumptions map[string]float64
	Note         string
}

func (c ConsumptionsManager) InsertConsumptions(args InsertConsumptionsArgs) (string, error) {
	child, err := c.CustomerStorage.GetChild(args.Code)
	if err != nil {
		return "", err
	}
	fmt.Println("Insertant consums de l'infant", child.Name, child.Surname)

	data, err := getInsertConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consumptions", viper.GetString("urls.hobbit"))
	return c.PostManager.PostPrint(url, data)
}

type ConsumptionBody struct {
	ProductID string  `json:"productId"`
	Units     float64 `json:"units"`
	Note      string  `json:"note"`
}
type ChildBody struct {
	Code         int               `json:"code"`
	Consumptions []ConsumptionBody `json:"consumptions"`
}
type ConsumptionsBody struct {
	YearMonth string      `json:"yearMonth"`
	Children  []ChildBody `json:"children"`
}

func buildConsumptionsBody(args InsertConsumptionsArgs) ConsumptionsBody {
	var consumptions []ConsumptionBody
	var note = args.Note
	for key, val := range args.Consumptions {
		consumptions = append(consumptions, ConsumptionBody{key, val, note})
		note = ""
	}
	consumptionsBody := ConsumptionsBody{
		YearMonth: viper.GetString("yearMonth"),
		Children: []ChildBody{
			{args.Code, consumptions},
		},
	}
	return consumptionsBody
}

func getInsertConsumptionsJson(args InsertConsumptionsArgs) ([]byte, error) {
	consumptionsBody := buildConsumptionsBody(args)
	bytes, err := json.Marshal(consumptionsBody)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}

func (c ConsumptionsManager) RectifyConsumptions(args InsertConsumptionsArgs) (string, error) {
	child, err := c.CustomerStorage.GetChild(args.Code)
	if err != nil {
		return "", err
	}
	fmt.Println("Rectificant els consums de l'infant", child.Name, child.Surname)

	data, err := getInsertConsumptionsJson(args)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/consumptions/rectification", viper.GetString("urls.hobbit"))
	return c.PostManager.PostPrint(url, data)
}

func (c ConsumptionsManager) BillConsumptions() error {
	fmt.Println("Facturant els consums pendents de facturar de tots els infants")
	return nil
}
