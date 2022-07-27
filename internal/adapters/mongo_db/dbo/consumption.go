package dbo

type Consumption struct {
	Id              string  `bson:"_id"`
	ChildId         int     `bson:"childCode"`
	ProductID       string  `bson:"productId"`
	Units           float64 `bson:"units"`
	YearMonth       string  `bson:"yearMonth"`
	Note            string  `bson:"note"`
	IsRectification bool    `bson:"isRectification"`
	InvoiceId       string  `bson:"invoiceId"`
}

func (c Consumption) GetId() interface{} {
	return c.Id
}
