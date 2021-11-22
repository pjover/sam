package translate

import (
	"strings"
	"time"
)

func WorkingDir(month time.Time) string {
	layout := "060100-Factures del mes January"
	dirName := month.Format(layout)
	englishMonth := month.Month().String()
	catalanMonth := MonthName(month)
	return strings.ReplaceAll(dirName, englishMonth, catalanMonth)
}

func MonthName(month time.Time) string {
	m := map[string]string{
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
		"December":  "de Decembre",
	}
	return m[month.Month().String()]
}
