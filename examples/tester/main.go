package main

import (
	"fmt"

	touchbar "github.com/LouisBrunner/go-touchbar/pkg"
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

func main() {
	tb := touchbar.New(barbuilder.Options{
		Configuration: barbuilder.Configuration{
			Items: []barbuilder.Item{
				// TODO: actually implement a useful example
				&barbuilder.Label{
					Content: &barbuilder.ContentImage{
						Image: barbuilder.TBAddTemplate,
					},
				},
				&barbuilder.Button{
					Title: "Hello",
					Image: barbuilder.SFSymbol("greaterthan.circle"),
					BezelColor: &barbuilder.RGBAColor{
						Red:   1.0,
						Green: 0.5,
						Blue:  0.2,
						Alpha: 1.0,
					},
					OnClick: func() {
						fmt.Printf("done\n")
					},
				},
				&barbuilder.Label{
					CommonProperties: barbuilder.CommonProperties{
						Principal: true,
					},
					Content: &barbuilder.ContentLabel{
						Text: "Hello",
						Color: &barbuilder.RGBAColor{
							Red:   0.3,
							Green: 1.0,
							Blue:  0.1,
							Alpha: 0.7,
						},
					},
				},
				&barbuilder.SpaceFlexible{},
				&barbuilder.Label{
					Content: &barbuilder.ContentImage{
						Image: barbuilder.SFSymbol("hammer"),
					},
				},
			},
		},
	})

	err := tb.Debug()
	if err != nil {
		panic(err)
	}

	err = tb.Uninstall()
	if err != nil {
		panic(err)
	}
}
