package mongo_express

import (
	"fmt"
	"github.com/pjover/sam/internal/domain/ports"
	"net/url"
)

type externalEditor struct {
	configService ports.ConfigService
	osService     ports.OsService
}

func NewExternalEditor(configService ports.ConfigService, osService ports.OsService) ports.ExternalEditor {
	return externalEditor{
		configService: configService,
		osService:     osService,
	}
}

func (e externalEditor) Edit(entity ports.Entity, id string) (string, error) {
	editUrl, err := e.getUrlPath(entity, id)
	if err != nil {
		return "", err
	}

	err = e.osService.OpenUrlInBrowser(editUrl)
	return editUrl, err
}

func (e externalEditor) getUrlPath(entity ports.Entity, id string) (string, error) {
	baseUrl := e.configService.GetString("urls.mongoExpress")
	switch entity {
	case ports.Customer:
		return fmt.Sprintf("%s/customer/%s", baseUrl, id), nil
	case ports.Invoice:
		return fmt.Sprintf("%s/invoice/%s", baseUrl, url.QueryEscape(fmt.Sprintf("\"%s\"", id))), nil
	case ports.Product:
		return fmt.Sprintf("%s/product/%s", baseUrl, url.QueryEscape(fmt.Sprintf("\"%s\"", id))), nil

	}
	return "", fmt.Errorf("unknown entity %#v", entity)
}
