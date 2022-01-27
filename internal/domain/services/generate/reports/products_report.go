package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/spf13/viper"
	"path"
	"sort"
	"strconv"
)

type ProductsReport struct {
	dbService ports.DbService
}

func NewProductsReport(dbService ports.DbService) ProductsReport {
	return ProductsReport{
		dbService: dbService,
	}
}

func (p ProductsReport) Run() (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("Generant l'informe de productes ...\n")

	products, err := p.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	contents := p.buildContents(products)

	filePath := path.Join(viper.GetString("dirs.reports"), viper.GetString("files.ProductsReport"))
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

	buffer.WriteString(fmt.Sprintf("Generat l'informe de productes a '%s'\n", filePath))
	return buffer.String(), nil
}

func (p ProductsReport) buildContents(products []model.Product) [][]string {
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

func (p ProductsReport) formatIsSubsidy(subsidy bool) string {
	if subsidy {
		return "Si"
	} else {
		return "No"
	}
}
