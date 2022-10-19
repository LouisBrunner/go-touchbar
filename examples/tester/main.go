package main

import (
	touchbar "github.com/LouisBrunner/go-touchbar/pkg"
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

func main() {
	tb := touchbar.New(barbuilder.Options{
		Configuration: barbuilder.Configuration{
			Items: []barbuilder.Item{
				// TODO: actually implement a useful example
				&barbuilder.Label{
					CommonProperties: barbuilder.CommonProperties{
						Principal: true,
					},
					Content: &barbuilder.ContentLabel{
						Text: "Hello",
					},
				},
				&barbuilder.Label{
					Content: &barbuilder.ContentImage{
						Image: barbuilder.SFSymbol("l.joystick.tilt.left.fill"),
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
