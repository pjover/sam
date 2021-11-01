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

func getInsertConsumptionsJson(args InsertConsumptionsArgs) ([]byte, error) {
	type Consumption struct {
		ProductID string  `json:"productId"`
		Units     float64 `json:"units"`
	}
	type Child struct {
		Code         int           `json:"code"`
		Consumptions []Consumption `json:"consumptions"`
	}
	type InsertConsumptionsJson struct {
		YearMonth string  `json:"yearMonth"`
		Children  []Child `json:"children"`
	}

	var consumptions []Consumption
	for key, val := range args.Consumptions {
		consumptions = append(consumptions, Consumption{key, val})
	}
	jsonText := InsertConsumptionsJson{
		YearMonth: viper.GetString("yearMonth"),
		Children: []Child{
			{args.Code, consumptions},
		},
	}

	bytes, err := json.Marshal(jsonText)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return bytes, nil
}
