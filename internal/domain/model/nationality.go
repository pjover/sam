package model

import (
	"fmt"
	"github.com/biter777/countries"
	"log"
)

type Nationality struct {
	countryCode countries.CountryCode
}

func NewNationality(alpha2Code string) (Nationality, error) {
	if alpha2Code == "" {
		return Nationality{}, nil
	}
	countryCode := countries.ByName(alpha2Code)
	if !countryCode.IsValid() {
		return Nationality{}, fmt.Errorf("'%s' is an invalid ISO 3166-1 alpha-2 country", alpha2Code)
	}
	return Nationality{countryCode: countryCode}, nil
}

func NewNationalityOrFatal(alpha2Code string) Nationality {
	nationality, err := NewNationality(alpha2Code)
	if err != nil {
		log.Fatal(err)
	}
	return nationality
}

func (n Nationality) String() string {
	if n.countryCode == countries.Unknown {
		return ""
	}
	return n.countryCode.Alpha2()
}
