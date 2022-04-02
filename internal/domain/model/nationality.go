package model

import (
	"fmt"
	"github.com/biter777/countries"
)

type Nationality struct {
	countryCode countries.CountryCode
}

var emptyNationality = Nationality{}

func (n Nationality) String() string {
	if n == emptyNationality {
		return ""
	}
	return n.countryCode.Alpha2()
}

func NewNationality(alpha2Code string) (Nationality, error) {
	countryCode := countries.ByName(alpha2Code)
	if !countryCode.IsValid() {
		return emptyNationality, fmt.Errorf("la nacionalitat '%s' no és vàlida", alpha2Code)
	}
	return Nationality{countryCode: countryCode}, nil
}

func NewNationalityOrEmpty(alpha2Code string) Nationality {
	newNationality, err := NewNationality(alpha2Code)
	if err != nil {
		return emptyNationality
	}
	return newNationality
}
