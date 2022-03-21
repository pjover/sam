package reports

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/pjover/sam/internal/domain/ports"
	"path"
)

type ReportService interface {
	SaveToFile(reportDefinition ReportDefinition, filePath string) error
}

type Report struct {
	configService ports.ConfigService
}

func NewReportService(configService ports.ConfigService) ReportService {
	return Report{
		configService: configService,
	}
}

type SubReport interface {
	GetTitle() string
	Render(maroto pdf.Maroto)
}

type ReportDefinition struct {
	// PageOrientation is a representation of a page orientation (either Portrait or Landscape)
	PageOrientation consts.Orientation
	// Title is the report's main title
	Title string
	// Footer is a text inserted after every report, if empty does not create ant footer
	Footer string
	// SubReports is a list of SubReports to include inside the main report, in order and one below the other
	SubReports []SubReport
}

func (r Report) SaveToFile(reportDefinition ReportDefinition, filePath string) error {
	maroto := r.setup(reportDefinition)
	r.header(maroto)
	r.footer(reportDefinition.Footer, maroto)
	r.title(reportDefinition.Title, maroto)
	r.subReports(reportDefinition.SubReports, maroto)
	return r.saveToFile(filePath, maroto)
}

func (r Report) setup(reportDefinition ReportDefinition) pdf.Maroto {
	m := pdf.NewMaroto(reportDefinition.PageOrientation, consts.A4)
	m.SetPageMargins(15, 10, 15)
	return m
}

func (r Report) header(maroto pdf.Maroto) {
	maroto.RegisterHeader(func() {
		maroto.Row(20, func() {
			maroto.Col(6, func() {
				_ = maroto.FileImage(
					path.Join(
						r.configService.GetString("dirs.config"),
						r.configService.GetString("files.logo"),
					),
					props.Rect{
						Left:    2,
						Center:  true,
						Percent: 80,
					})
			})

			maroto.ColSpace(2)

			maroto.Col(4, func() {
				maroto.Text(r.configService.GetString("business.name"), props.Text{
					Style:       consts.BoldItalic,
					Size:        8,
					Align:       consts.Left,
					Extrapolate: false,
				})
				maroto.Text(r.configService.GetString("business.addressLine1"), props.Text{
					Top:   3,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(r.configService.GetString("business.addressLine2"), props.Text{
					Top:   6,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(r.configService.GetString("business.addressLine3"), props.Text{
					Top:   9,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(r.configService.GetString("business.addressLine4"), props.Text{
					Top:   12,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(r.configService.GetString("business.taxIdLine"), props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
			})
		})
	})
}

func (r Report) footer(footer string, maroto pdf.Maroto) {
	if footer == "" {
		return
	}
	maroto.RegisterFooter(func() {
		maroto.Row(4, func() {
			maroto.Col(12, func() {
				maroto.Text(
					footer,
					props.Text{
						Top:   4,
						Style: consts.Italic,
						Size:  8,
						Align: consts.Right,
					})
			})
		})
	})
}

func (r Report) title(title string, maroto pdf.Maroto) {
	maroto.Row(20, func() {
		maroto.Col(12, func() {
			maroto.Text(
				title,
				props.Text{
					Top:   4,
					Style: consts.Bold,
					Align: consts.Center,
					Color: color.Color{
						Red:   0,
						Green: 51,
						Blue:  51,
					},
					Size: 18,
				})
		})
	})
}

func (r Report) subReports(subReports []SubReport, maroto pdf.Maroto) {
	for _, subReport := range subReports {
		r.subTitle(subReport.GetTitle(), maroto)
		subReport.Render(maroto)
	}
}

func (r Report) subTitle(subTitle string, maroto pdf.Maroto) {
	if subTitle == "" {
		return
	}

	maroto.Row(20, func() {
		maroto.Col(12, func() {
			maroto.Text(
				subTitle,
				props.Text{
					Top:   8,
					Style: consts.Bold,
					Align: consts.Left,
					Color: color.Color{
						Red:   0,
						Green: 51,
						Blue:  51,
					},
					Size: 14,
				})
		})
	})
}

func (r Report) saveToFile(filePath string, maroto pdf.Maroto) error {
	err := maroto.OutputFileAndClose(filePath)
	if err != nil {
		return fmt.Errorf("error while saving report to file '%s': %s", filePath, err)
	}
	return nil
}
