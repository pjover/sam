package model

import "strings"

type Language uint

const (
	Invalid Language = iota
	Catalan
	English
	Spanish
)

var formatValues = []string{
	"Indefinit",
	"Català",
	"Anglès",
	"Espanyol",
}

func (p Language) Format() string {
	return formatValues[p]
}

var stringValues = []string{
	"",
	"CA",
	"EN",
	"ES",
}

func (p Language) String() string {
	return stringValues[p]
}

func NewLanguage(value string) Language {

	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return Language(i)
		}
	}
	return Invalid
}
