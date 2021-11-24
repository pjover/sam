package consum

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
)

type ConsumptionBody struct {
	ProductID string  `json:"productId"`
	Units     float64 `json:"units"`
	Note      string  `json:"note"`
}
type ChildBody struct {
	Code         int               `json:"code"`
	Consumptions []ConsumptionBody `json:"consumptions"`
}
type Body struct {
	YearMonth string      `json:"yearMonth"`
	Children  []ChildBody `json:"children"`
}

func getConsumptionsJson(args CustomerConsumptionsArgs) ([]byte, error) {
	consumptionsBody := buildConsumptionsBody(args)
	bytes, err := json.Marshal(consumptionsBody)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(bytes[:]))
	return bytes, nil
}

func buildConsumptionsBody(args CustomerConsumptionsArgs) Body {
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
