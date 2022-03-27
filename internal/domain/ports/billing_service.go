package ports

type BillingService interface {
	InsertConsumptions(id int, consumptions map[string]float64, note string) (string, error)
	BillConsumptions() (string, error)
	RectifyConsumptions(id int, consumptions map[string]float64, note string) (string, error)
}
