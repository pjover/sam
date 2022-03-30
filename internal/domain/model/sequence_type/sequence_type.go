package sequence_type

import "strings"

type SequenceType uint

const (
	Invalid              SequenceType = iota
	StandardInvoice                   // Standard invoices with bank direct debit payment type
	SpecialInvoice                    // Special invoices, with other payments types
	RectificationInvoice              // Rectification invoices, correct standard or special invoice
	Customer                          // Customers
)

var stringValues = []string{
	"",
	"STANDARD_INVOICE",
	"SPECIAL_INVOICE",
	"RECTIFICATION_INVOICE",
	"CUSTOMER",
}

func (p SequenceType) String() string {
	return stringValues[p]
}

func NewSequenceType(value string) SequenceType {
	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return SequenceType(i)
		}
	}
	return Invalid
}

var formatValues = []string{
	"Indefinit",
	"Factura (rebut)",
	"Factura (no rebut)",
	"Rectificació",
	"Client",
}

func (s SequenceType) Format() string {
	return formatValues[s]
}

func (s SequenceType) Prefix() string {
	switch s {
	case StandardInvoice:
		return "F"
	case SpecialInvoice:
		return "X"
	case RectificationInvoice:
		return "R"
	case Customer:
		return "C"
	default:
		return ""
	}
}
