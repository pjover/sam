package catalan

import (
	"strings"
	"time"
)

func WorkingDir(workingTime time.Time) string {
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
	layout := "060100-Factures del mes January"
	dirName := workingTime.Format(layout)
	englishMonth := workingTime.Month().String()
	catalanMonth := m[englishMonth]
	return strings.ReplaceAll(dirName, englishMonth, catalanMonth)
}
