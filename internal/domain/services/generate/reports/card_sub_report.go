package reports

import (
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type CardSubReport struct {
	// Title is a text that precedes the table
	Title string
	// Alignment is the alignment of the text (header and content) inside the columns
	Align consts.Align // TODO move Align to each column if possible with text styles
	// Captions is a list of captions for every table's colum or card file
	Captions []string
	// Widths is a list of columns' widths
	Widths []uint
	// Data is a list of lists containing every cell text
	Data [][]string
}

func (c CardSubReport) GetTitle() string {
	return c.Title
}

func (c CardSubReport) Render(maroto pdf.Maroto) {
	var captionProps = props.Text{
		Style: consts.BoldItalic,
		Top:   0,
		Size:  10,
		Align: consts.Left,
	}
	var cellProps = props.Text{
		Top:             0,
		VerticalPadding: 2,
		Size:            10,
		Align:           consts.Left,
	}
	for row, caption := range c.Captions {
		maroto.Row(4, func() {
			maroto.Col(c.Widths[0], func() {
				maroto.Text(caption, captionProps)
			})
			for col, cell := range c.Data[row] {
				maroto.Col(c.Widths[col+1], func() {
					maroto.Text(cell, cellProps)
				})
			}
		})
	}
}
