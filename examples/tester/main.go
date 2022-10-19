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
					Title:      "Hello",
					Image:      barbuilder.SFSymbol("greaterthan.circle"),
					BezelColor: barbuilder.Color("blue"),
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
