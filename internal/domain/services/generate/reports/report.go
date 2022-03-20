package reports

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/spf13/viper"
	"log"
	"path"
	"time"
)

type Style uint

const (
	Table Style = iota
	Card
)

type SubReport struct {
	// Style is the report style (Table or Card)
	Style Style
	// Alignment is the alignment of the text (header and content) inside the columns
	Align consts.Align
	// Captions is a list of captions for every table's colum or card file
	Captions []string
	// Widths is a list of columns' widths
	Widths []uint
	// Data is a list of lists containing every cell text
	Data [][]string
}

type ReportDefinition struct {
	// PageOrientation is a representation of a page orientation (either Portrait or Landscape)
	PageOrientation consts.Orientation
	// Title is the report's main title
	Title string
	// FilePath is where the report will be saved
	FilePath string
	// SubReports is a list of SubReports to include inside the main report, in order one below the other
	SubReports []SubReport
}

func (r ReportDefinition) Generate() error {
	maroto := r.setup()
	r.header(maroto)
	r.footer(maroto)
	r.title(maroto)
	r.body(maroto)
	return r.save(maroto)
}

func (r ReportDefinition) setup() pdf.Maroto {
	m := pdf.NewMaroto(r.PageOrientation, consts.A4)
	m.SetPageMargins(15, 10, 15)
	return m
}

func (r ReportDefinition) header(maroto pdf.Maroto) {
	maroto.RegisterHeader(func() {
		maroto.Row(20, func() {
			maroto.Col(6, func() {
				_ = maroto.FileImage(
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

			maroto.ColSpace(2)

			maroto.Col(4, func() {
				maroto.Text(viper.GetString("business.name"), props.Text{
					Style:       consts.BoldItalic,
					Size:        8,
					Align:       consts.Left,
					Extrapolate: false,
				})
				maroto.Text(viper.GetString("business.addressLine1"), props.Text{
					Top:   3,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(viper.GetString("business.addressLine2"), props.Text{
					Top:   6,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(viper.GetString("business.addressLine3"), props.Text{
					Top:   9,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(viper.GetString("business.addressLine4"), props.Text{
					Top:   12,
					Size:  8,
					Align: consts.Left,
				})
				maroto.Text(viper.GetString("business.taxIdLine"), props.Text{
					Top:   15,
					Style: consts.BoldItalic,
					Size:  8,
					Align: consts.Left,
				})
			})
		})
	})
}

func (r ReportDefinition) footer(maroto pdf.Maroto) {
	maroto.RegisterFooter(func() {
		maroto.Row(4, func() {
			maroto.Col(12, func() {
				maroto.Text(time.Now().Format("2006-01-02"),
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

func (r ReportDefinition) title(maroto pdf.Maroto) {
	maroto.Row(20, func() {
		maroto.Col(12, func() {
			maroto.Text(
				r.Title,
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

func (r ReportDefinition) body(maroto pdf.Maroto) {
	for _, subReport := range r.SubReports {
		r.subReport(maroto, subReport)
	}
}

func (r ReportDefinition) subReport(maroto pdf.Maroto, subReport SubReport) {
	switch subReport.Style {
	case Table:
		r.table(maroto, subReport)
	case Card:
		r.card(maroto, subReport)
	}
}

func (r ReportDefinition) table(maroto pdf.Maroto, subReport SubReport) {

	maroto.TableList(subReport.Captions, subReport.Data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: subReport.Widths,
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: subReport.Widths,
		},
		Align: subReport.Align,
		AlternatedBackground: &color.Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		},
		HeaderContentSpace: 1,
		Line:               false,
	})
}

func (r ReportDefinition) card(maroto pdf.Maroto, report SubReport) {

}

func (r ReportDefinition) save(maroto pdf.Maroto) error {
	err := maroto.OutputFileAndClose(r.FilePath)
	if err != nil {
		return fmt.Errorf("error while saving report to file '%s': %s", r.FilePath, err)
	}
	return nil
}

type Column struct {
	Name  string
	Width uint
}

type ReportInfo struct {
	Orientation consts.Orientation
	Align       consts.Align
	Title       string
	FilePath    string
	Columns     []Column
	Contents    [][]string
}

func Report(reportInfo ReportInfo) error {
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
