package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	touchbar "github.com/LouisBrunner/go-touchbar/pkg"
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

var spinValues = []string{"🍒", "💎", "7️⃣", "🍊", "🔔", "⭐", "🍇", "🍀"}

func randomChoice[T any](choices []T) (*T, error) {
	v, err := rand.Int(rand.Reader, big.NewInt(int64(len(choices))))
	if err != nil {
		return nil, err
	}
	return &choices[int(v.Uint64())], nil
}

func getRandomValue() (*string, error) {
	return randomChoice(spinValues)
}

func main() {
	tb := touchbar.New(barbuilder.Options{})
	var items, catalogItems, startItems []barbuilder.Item
	updater := func() {
		err := tb.Update(barbuilder.Configuration{Items: items})
		if err != nil {
			fmt.Printf("error: %+v\n", err)
		}
	}

	startItems = []barbuilder.Item{
		&barbuilder.Label{
			Content: &barbuilder.ContentLabel{
				Text: "Go Touch Bar",
			},
		},
		&barbuilder.SpaceLarge{},
		makeDemo(updater),
		&barbuilder.SpaceSmall{},
		&barbuilder.Button{
			Title: "Catalog",
			Image: barbuilder.TBBookmarksTemplate,
			OnClick: func() {
				items = catalogItems
				updater()
			},
		},
	}

	catalogItems = []barbuilder.Item{
		&barbuilder.Button{
			Title: "Close",
			OnClick: func() {
				items = startItems
				updater()
			},
		},
		&barbuilder.Label{
			Content: &barbuilder.ContentLabel{
				Text: "Catalog",
			},
		},
		makeLabelCatalog(),
		// makeButtonCatalog(updater),
		// makePopoverCatalog(),
	}

	items = catalogItems
	err := tb.Debug(barbuilder.Configuration{Items: items})
	if err != nil {
		panic(err)
	}

	err = tb.Uninstall()
	if err != nil {
		panic(err)
	}
}
