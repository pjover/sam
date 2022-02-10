package dbo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
)

func Decimal128ToFloat64(d128 primitive.Decimal128) float64 {
	float, err := strconv.ParseFloat(d128.String(), 64)
	if err != nil {
		panic(fmt.Sprintf("converting Decimal128 %s to float", d128))
	}
	return float
}

func Float64ToDecimal128(f64 float64) primitive.Decimal128 {
	d128, err := primitive.ParseDecimal128(fmt.Sprintf("%f", f64))
	if err != nil {
		panic(fmt.Sprintf("converting float %f to Decimal128", f64))
	}
	return d128
}
