package model

import (
	"fmt"
	validator "github.com/criptalia/spanish_dni_validator"
)

type TaxId struct {
	id string
}

var emptyTaxId = TaxId{}

func (t TaxId) String() string {
	if t == emptyTaxId {
		return ""
	}
	return t.id
}

func NewTaxId(taxId string) (TaxId, error) {
	if validator.IsValid(taxId) {
		return TaxId{id: taxId}, nil
	}
	return emptyTaxId, fmt.Errorf("el DNI/NIE/CIF '%s' no és vàlid", taxId)
}

func NewTaxIdOrEmpty(taxId string) TaxId {
	newTaxId, err := NewTaxId(taxId)
	if err != nil {
		return emptyTaxId
	}
	return newTaxId
}
