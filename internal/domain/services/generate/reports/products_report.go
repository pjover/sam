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
	"time"
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

	report := Report{
		PageOrientation: consts.Portrait,
		Title:           "Llistat de productes",
		Footer:          time.Now().Format("2006-01-02"),
		SubReports: []SubReport{
			TableSubReport{
				Align: consts.Left,
				Captions: []string{
					"Codi",
					"Nom",
					"Preu",
					"IVA",
					"És ajuda?",
				},
				Widths: []uint{
					1,
					7,
					2,
					1,
					1,
				},
				Data: p.buildContents(products),
			},
		},
	}

	filePath := path.Join(
		viper.GetString("dirs.reports"),
		viper.GetString("files.ProductsReport"),
	)
	err = report.SaveToFile(filePath)
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("Generat l'informe de productes a '%s'\n", filePath))
	return buffer.String(), nil
}

func (p ProductsReport) buildContents(products []model.Product) [][]string {
	var data [][]string
	for _, product := range products {
		var line = []string{
			product.Id,
			product.Name,
			strconv.FormatFloat(product.Price, 'f', 2, 64),
			strconv.FormatFloat(product.TaxPercentage, 'f', 2, 64),
			p.formatIsSubsidy(product.IsSubsidy),
		}
		data = append(data, line)
	}
	sort.SliceStable(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	return data
}

func (p ProductsReport) formatIsSubsidy(subsidy bool) string {
	if subsidy {
		return "Si"
	} else {
		return "No"
	}
}
