package main

import (
	touchbar "github.com/LouisBrunner/go-touchbar/pkg"
)

func main() {
	tb := touchbar.New(touchbar.Options{
		Configuration: touchbar.Configuration{
			Items: touchbar.Items{
				// TODO: actually implement a useful example
				&touchbar.Label{},
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
