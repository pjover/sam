package payment_type

type PaymentType uint

const (
	Invalid PaymentType = iota
	BankDirectDebit
	BankTransfer
	Voucher
	Cash
	Rectification
)

var values = []string{
	"Indefinit",
	"Rebut",
	"Tranferència",
	"Xec escoleta",
	"Efectiu",
	"Rectificació",
}

func (s PaymentType) String() string {
	return values[s]
}
