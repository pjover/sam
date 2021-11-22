package reports

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/spf13/viper"
	"log"
	"path"
	"time"
)

type Column struct {
	Name  string
	Width uint
}

type ReportInfo struct {
	Orientation consts.Orientation
	Align       consts.Align
	Title       string
	Columns     []Column
	Contents    [][]string
	FilePath    string
}

func PdfReport(reportInfo ReportInfo) error {
	m := setupStandardPage(reportInfo)
	header(m)
	footer(m)
	title(m, reportInfo)
	table(m, reportInfo)
	err := m.OutputFileAndClose(reportInfo.FilePath)
	if err != nil {
		return err
	}
	return nil
}

func setupStandardPage(reportInfo ReportInfo) pdf.Maroto {
	m := pdf.NewMaroto(reportInfo.Orientation, consts.A4)
	m.SetPageMargins(15, 10, 15)
	return m
}

func header(m pdf.Maroto) {
	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(6, func() {
				_ = m.FileImage(
					path.Join(
						viper.GetString("dirs.config"),
						viper.GetString("files.logo"),
					),
					props.Rect{
						Left:    2,
						Center:  true,
						Percent: 80,
					})
			})

			m.ColSpace(2)

			m.Col(4, func() {
				m.Text(viper.GetString("business.name"), props.Text{
					Style:       consts.BoldItalic,
					Size:        8,
					Align:       consts.Left,
					Extrapolate: false,
				})
				m.Text(viper.GetString("business.addressLine1"), props.Text{
					Top:   3,
					Size:  8,
					Align: consts.Left,
				})
				m.Text(viper.GetString("business.addressLine2"), props.Text{
					Top:   6,
					Size:  8,
					Align: consts.Left,
				})
				m.Text(viper.GetString("business.addressLine3"), props.Text{
					Top:   9,
					Size:  8,
					Align: consts.Left,
				})
				m.Text(viper.GetString("business.addressLine4"), props.Text{
					Top:   12,
					Size:  8,
					Align: consts.Left,
				})
				m.Text(viper.GetString("business.taxIdLine"), props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
			})
		})
	})
}

func footer(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(4, func() {
			m.Col(12, func() {
				m.Text(time.Now().Format("2006-01-02"),
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

func title(m pdf.Maroto, reportInfo ReportInfo) {
	m.Row(20, func() {
		m.Col(12, func() {
			m.Text(
				reportInfo.Title,
				props.Text{
					Top:   4,
					Style: consts.Bold,
					Align: consts.Center,
					Color: color.Color{
						Red:   0,
						Green: 51,
						Blue:  51,
					},
					Size: 24,
				})
		})
	})
}

func table(m pdf.Maroto, reportInfo ReportInfo) {
	headers, widths := extractFromColumns(reportInfo.Columns)

	m.TableList(headers, reportInfo.Contents, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: widths,
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: widths,
		},
		Align: reportInfo.Align,
		AlternatedBackground: &color.Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		},
		HeaderContentSpace: 1,
		Line:               false,
	})
}

func extractFromColumns(columns []Column) ([]string, []uint) {
	var headers []string
	var widths []uint
	var widthsSum uint
	for _, column := range columns {
		widthsSum += column.Width
		headers = append(headers, column.Name)
		widths = append(widths, column.Width)
	}
	if widthsSum != 12 {
		log.Fatal("Columns widths always must sum 12")
	}
	return headers, widths
}
