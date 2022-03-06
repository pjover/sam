package dbo

type Sequence struct {
	Id      string `bson:"_id"`
	Counter int    `bson:"counter"`
}

func (s Sequence) GetId() interface{} {
	return s.Id
}
