package model

import (
	"fmt"
	taxId "github.com/criptalia/spanish_dni_validator"
	"log"
)

type TaxId struct {
	id string
}

func NewTaxId(id string) (TaxId, error) {
	if id == "" {
		return TaxId{id: id}, nil
	}
	if taxId.IsValid(id) {
		return TaxId{id: id}, nil
	}
	return TaxId{}, fmt.Errorf("el DNI/NIE/CIF '%s' no és vàlid", id)
}

func NewTaxIdOrFatal(id string) TaxId {
	taxId, err := NewTaxId(id)
	if err != nil {
		log.Fatal(err)
	}
	return taxId
}

func (t TaxId) String() string {
	return t.id
}
