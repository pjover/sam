package dbo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func Decimal128ToFloat64(d128 primitive.Decimal128) float64 {
	float, err := strconv.ParseFloat(d128.String(), 64)
	if err != nil {
		panic(fmt.Sprintf("converting Decimal128 %s to float", d128.String()))
	}
	return float
}
