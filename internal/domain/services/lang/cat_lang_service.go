package lang

import (
	"strings"
	"time"
)

type catLangService struct {
}

func NewCatLangService() LangService {
	return catLangService{}
}

var layout = "060100-Factures del mes January"

func (c catLangService) WorkingDir(workingTime time.Time) string {
	dirName := workingTime.Format(layout)
	englishMonth := workingTime.Month().String()
	catalanMonth := c.MonthName(workingTime.Month())
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

func (c catLangService) MonthName(month time.Month) string {
	return m[month.String()]
}
