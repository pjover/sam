package lang

import (
	"log"
	"time"
)

type LangService interface {
	WorkingDir(month time.Time) string
	MonthName(month time.Time) string
}

func NewLangService(language string) LangService {
	switch language {
	case "cat":
	case "ca_ES.UTF-8":
		return catLangService{}
	default:
		log.Fatalf("Language '%s' not implemented", language)
	}
	return nil
}
