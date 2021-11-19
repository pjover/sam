package generate

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/spf13/viper"
	"path"
	"sam/internal/util"
	"sam/model"
	"sort"
	"strings"
)

type ProductsReportGenerator struct {
	getManager util.HttpGetManager
}

func NewProductsReportGenerator(getManager util.HttpGetManager) ProductsReportGenerator {
	return ProductsReportGenerator{
		getManager,
	}
}

func (p ProductsReportGenerator) generate() (string, error) {
	products, err := p.getProducts()
	if err != nil {
		return "", err
	}

	contents := p.buildContents(products)

	filePath := path.Join(viper.GetString("dirs.reports"), viper.GetString("files.productsReport"))
	reportInfo := util.ReportInfo{
		consts.Portrait,
		consts.Left,
		"Llistat de productes",
		[]util.Column{
			{"Codi", 2},
			{"Nom", 4},
			{"Preu", 2},
			{"IVA", 2},
			{"És ajuda?", 2},
		},
		contents,
		filePath,
	}
	err = util.PdfReport(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de productes a '%s'", filePath), nil
}

type products struct {
	Embedded struct {
		Products []model.Product `json:"products"`
	} `json:"_embedded"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Profile struct {
			Href string `json:"href"`
		} `json:"profile"`
	} `json:"_links"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

func (p ProductsReportGenerator) getProducts() (*products, error) {
	url := fmt.Sprintf(
		"%s/products?page=0&size=999",
		viper.GetString("urls.hobbit"),
	)
	products := new(products)
	err := p.getManager.Type(url, products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductsReportGenerator) buildContents(products *products) [][]string {
	var contents [][]string
	for _, product := range products.Embedded.Products {
		var line = []string{
			p.getCode(product),
			product.Name,
			fmt.Sprintf("%.2f", product.Price),
			fmt.Sprintf("%.2f", product.TaxPercentage),
			p.formatIsSubsidy(product.IsSubsidy),
		}
		contents = append(contents, line)
	}
	sort.SliceStable(contents, func(i, j int) bool {
		return contents[i][0] < contents[j][0]
	})
	return contents
}

func (p ProductsReportGenerator) getCode(product model.Product) string {
	url := product.Links.Self.Href
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func (p ProductsReportGenerator) formatIsSubsidy(subsidy bool) string {
	if subsidy {
		return "Si"
	} else {
		return "No"
	}
}