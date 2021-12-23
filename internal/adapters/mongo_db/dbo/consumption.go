package dbo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Consumption struct {
	Code            string               `bson:"_id"`
	ChildCode       int                  `bson:"childCode"`
	ProductID       string               `bson:"productId"`
	Units           primitive.Decimal128 `bson:"units"`
	YearMonth       string               `bson:"yearMonth"`
	IsRectification bool                 `bson:"isRectification"`
	InvoiceCode     string               `bson:"invoiceId"`
}
