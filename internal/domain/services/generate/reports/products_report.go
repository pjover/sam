package reports

import (
	"bytes"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
	"sort"
	"strconv"
)

type ProductsReport struct {
	configService ports.ConfigService
	dbService     ports.DbService
	osService     ports.OsService
}

func NewProductsReport(configService ports.ConfigService, dbService ports.DbService, osService ports.OsService) ProductsReport {
	return ProductsReport{
		configService: configService,
		dbService:     dbService,
		osService:     osService,
	}
}

func (p ProductsReport) Run() (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString("Generant l'informe de productes ...\n")

	products, err := p.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	reportDefinition := ReportDefinition{
		PageOrientation: consts.Portrait,
		Title:           "Llistat de productes",
		Footer:          p.osService.Now().Format(domain.YearMonthDayLayout),
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

	reportsDir := p.configService.GetReportsDirectory()
	filePath := path.Join(
		reportsDir,
		p.configService.GetString("files.ProductsReport"),
	)

	reportService := NewReportService(p.configService)
	err = reportService.SaveToFile(reportDefinition, filePath)
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
