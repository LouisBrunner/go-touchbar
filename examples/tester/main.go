package main

import (
	"crypto/rand"
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
	var items []barbuilder.Item
	updater := func() {
		tb.Update(barbuilder.Configuration{Items: items})
	}

	items = []barbuilder.Item{
		&barbuilder.Label{
			Content: &barbuilder.ContentLabel{
				Text: "Go Touch Bar",
			},
		},
		&barbuilder.SpaceLarge{},
		makeDemo(updater),
		&barbuilder.SpaceSmall{},
		&barbuilder.Popover{
			CollapsedText:  "Catalog",
			CollapsedImage: barbuilder.TBBookmarksTemplate,
			Bar: []barbuilder.Item{
				&barbuilder.Label{
					Content: &barbuilder.ContentLabel{
						Text: "Catalog",
					},
				},
			},
		},
	}

	err := tb.Debug(barbuilder.Configuration{Items: items})
	if err != nil {
		panic(err)
	}

	err = tb.Uninstall()
	if err != nil {
		panic(err)
	}
}
