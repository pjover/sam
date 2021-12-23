package dbo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	Id            string               `bson:"_id"`
	Name          string               `bson:"name"`
	ShortName     string               `bson:"shortName"`
	Price         primitive.Decimal128 `bson:"price"`
	TaxPercentage primitive.Decimal128 `bson:"taxPercentage"`
	IsSubsidy     bool                 `bson:"isSubsidy"`
}
