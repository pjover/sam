package model

import "strings"

type Language uint

const (
	Undefined Language = iota
	Catalan
	English
	Spanish
)

var stringValues = []string{
	"",
	"CA",
	"EN",
	"ES",
}

func (l Language) String() string {
	return stringValues[l]
}

func NewLanguage(value string) Language {

	value = strings.ToLower(value)
	for i, val := range stringValues {
		if strings.ToLower(val) == value {
			return Language(i)
		}
	}
	return Undefined
}

var formatValues = []string{
	"Indefinit",
	"Català",
	"Anglès",
	"Espanyol",
}

func (l Language) Format() string {
	return formatValues[l]
}
