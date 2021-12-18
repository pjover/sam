package reports

import (
	"fmt"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/generate"
	"github.com/spf13/viper"
)

type ProductsReportGenerator struct {
	dbService ports.DbService
}

func NewProductsReportGenerator(dbService ports.DbService) generate.Generator {
	return ProductsReportGenerator{
		dbService: dbService,
	}
}

func (p ProductsReportGenerator) Generate() (string, error) {
	fmt.Println("Generant l'informe de productes ...")

	products, err := p.dbService.GetAllProducts()
	if err != nil {
		return "", err
	}

	contents := p.buildContents(products)

	filePath := path.Join(viper.GetString("dirs.reports"), viper.GetString("files.productsReport"))
	reportInfo := ReportInfo{
		consts.Portrait,
		consts.Left,
		"Llistat de productes",
		[]Column{
			{"Codi", 2},
			{"Nom", 4},
			{"Preu", 2},
			{"IVA", 2},
			{"És ajuda?", 2},
		},
		contents,
		filePath,
	}
	err = Report(reportInfo)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Generat l'informe de productes a '%s'", filePath), nil
}

func (p ProductsReportGenerator) buildContents(products []model.Product) [][]string {
	var contents [][]string
	for _, product := range products {
		var line = []string{
			product.Id,
			product.Name,
			strconv.FormatFloat(product.Price, 'f', 2, 64),
			strconv.FormatFloat(product.TaxPercentage, 'f', 2, 64),
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
