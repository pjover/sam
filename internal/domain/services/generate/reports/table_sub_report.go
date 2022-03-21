package reports

import (
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type TableSubReport struct {
	// Title is a text that precedes the table
	Title string
	// Alignment is the alignment of the text (header and content) inside the columns
	Align consts.Align
	// Captions is a list of captions for every table's colum or card file
	Captions []string
	// Widths is a list of columns' widths
	Widths []uint
	// Data is a list of lists containing every cell text
	Data [][]string
}

func (r TableSubReport) GetTitle() string {
	return r.Title
}

func (r TableSubReport) Render(maroto pdf.Maroto) {
	maroto.TableList(r.Captions, r.Data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      9,
			GridSizes: r.Widths,
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: r.Widths,
		},
		Align: r.Align,
		AlternatedBackground: &color.Color{
			Red:   200,
			Green: 200,
			Blue:  200,
		},
		HeaderContentSpace: 1,
		Line:               false,
	})
}
