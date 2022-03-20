package reports

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/spf13/viper"
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

type Report struct {
	// PageOrientation is a representation of a page orientation (either Portrait or Landscape)
	PageOrientation consts.Orientation
	// Title is the report's main title
	Title string
	// SubReports is a list of SubReports to include inside the main report, in order and one below the other
	SubReports []SubReport
}

func (r Report) SaveToFile(filePath string) error {
	maroto := r.setup()
	r.header(maroto)
	r.footer(maroto)
	r.title(maroto)
	r.body(maroto)
	return r.saveToFile(maroto, filePath)
}

func (r Report) setup() pdf.Maroto {
	m := pdf.NewMaroto(r.PageOrientation, consts.A4)
	m.SetPageMargins(15, 10, 15)
	return m
}

func (r Report) header(maroto pdf.Maroto) {
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

func (r Report) footer(maroto pdf.Maroto) {
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

func (r Report) title(maroto pdf.Maroto) {
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

func (r Report) body(maroto pdf.Maroto) {
	for _, subReport := range r.SubReports {
		r.subReport(maroto, subReport)
	}
}

func (r Report) subReport(maroto pdf.Maroto, subReport SubReport) {
	switch subReport.Style {
	case Table:
		r.table(maroto, subReport)
	case Card:
		r.card(maroto, subReport)
	}
}

func (r Report) table(maroto pdf.Maroto, subReport SubReport) {

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

func (r Report) card(maroto pdf.Maroto, report SubReport) {

}

func (r Report) saveToFile(maroto pdf.Maroto, filePath string) error {
	err := maroto.OutputFileAndClose(filePath)
	if err != nil {
		return fmt.Errorf("error while saving report to file '%s': %s", filePath, err)
	}
	return nil
}
