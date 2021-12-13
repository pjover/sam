package mongo_express

import (
	"fmt"
	"github.com/pjover/sam/internal/core/ports"
	"net/url"
)

type externalEditor struct {
	cfgService ports.ConfigService
	osService  ports.OsService
}

func NewExternalEditor(cfgService ports.ConfigService, osService ports.OsService) ports.ExternalEditor {
	return externalEditor{
		cfgService: cfgService,
		osService:  osService,
	}
}

func (e externalEditor) Edit(entity ports.Entity, code string) (string, error) {
	editUrl, err := e.getUrlPath(entity, code)
	if err != nil {
		return "", err
	}

	err = e.osService.OpenUrlInBrowser(editUrl)
	return editUrl, err
}

func (e externalEditor) getUrlPath(entity ports.Entity, code string) (string, error) {
	baseUrl := e.cfgService.Get("urls.mongoExpress")
	switch entity {
	case ports.Customer:
		return fmt.Sprintf("%s/customer/%s", baseUrl, code), nil
	case ports.Invoice:
		return fmt.Sprintf("%s/invoice/%s", baseUrl, url.QueryEscape(fmt.Sprintf("\"%s\"", code))), nil
	case ports.Product:
		return fmt.Sprintf("%s/product/%s", baseUrl, url.QueryEscape(fmt.Sprintf("\"%s\"", code))), nil

	}
	return "", fmt.Errorf("unknown entity %#v", entity)
}
