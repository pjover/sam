package lang

import (
	"strings"
	"time"
)

type catLangService struct {
}

var layout = "060100-Factures del mes January"

func (c catLangService) WorkingDir(month time.Time) string {
	dirName := month.Format(layout)
	englishMonth := month.Month().String()
	catalanMonth := c.MonthName(month)
	return strings.ReplaceAll(dirName, englishMonth, catalanMonth)
}

var m = map[string]string{
	"January":   "de Gener",
	"February":  "de Febrer",
	"March":     "de Març",
	"April":     "d'Abril",
	"May":       "de Maig",
	"June":      "de Juny",
	"July":      "de Juliol",
	"August":    "d'Agost",
	"September": "de Setembre",
	"October":   "d'Octubre",
	"November":  "de Novembre",
	"December":  "de Desembre",
}

func (c catLangService) MonthName(month time.Time) string {
	return m[month.Month().String()]
}
