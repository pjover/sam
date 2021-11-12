package adm

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/spf13/viper"
	"log"
	"os"
	"path"
)

func CustomerReportPdf(filePath string) error {
	m := SetupStandardPage()

	err := header(m)
	if err != nil {
		return err
	}

	err = m.OutputFileAndClose(filePath)
	if err != nil {
		return err
	}
	return nil
}

func SetupStandardPage() pdf.Maroto {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)
	return m
}

func header(m pdf.Maroto) error {
	logo := path.Join(viper.GetString("dirs.config"), viper.GetString("files.logo"))
	if _, err := os.Stat(logo); err != nil {
		return fmt.Errorf("No s'ha trobat el fitxer del logo '%s'", logo)
	}
	m.RegisterHeader(func() {
		m.Row(40, func() {
			m.Col(12, func() {
				err := m.FileImage(logo,
					props.Rect{
						Center:  true,
						Percent: 90,
					})
				if err != nil {
					log.Fatal(err)
				}
			})
		})
	})
	m.Row(10, func() {
		m.Col(12, func() {
			m.Text("Prepared for you by the Div Rhino Fruit Company", props.Text{
				Top:   4,
				Style: consts.Bold,
				Align: consts.Center,
				Color: color.Color{
					Red:   0,
					Green: 51,
					Blue:  51,
				},
				Size: 30,
			})
		})
	})
	return nil
}
