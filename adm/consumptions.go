package adm

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"sam/comm"
	"sam/storage"
)

type InsertConsumptionsArgs struct {
	Code         int
	Consumptions map[string]float64
	Note         string
}

func InsertConsumptions(args InsertConsumptionsArgs) error {
	child, err := storage.GetChild(args.Code)
	if err != nil {
		return err
	}
	fmt.Println("Insertant consums de l'infant", child.Name, child.Surname)

	data, err := getInsertConsumptionsJson(args)
	if err != nil {
		return err
	}

	err = comm.PrintPost("http://localhost:8080/consumptions", data)
	if err != nil {
		return err
	}
	return nil
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
