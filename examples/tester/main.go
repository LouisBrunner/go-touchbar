package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	touchbar "github.com/LouisBrunner/go-touchbar/pkg"
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
	"github.com/LouisBrunner/go-touchbar/pkg/barutils"
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

	makeUpdater := func(switcher barutils.Switcher) func() {
		return func() {
			err := switcher.Update()
			if err != nil {
				fmt.Printf("could not update: %v\n", err)
			}
		}
	}

	config := barutils.MakeStackableBar(tb, func(switcher barutils.Switcher) []barbuilder.Item {
		update := makeUpdater(switcher)
		return []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Go Touch Bar",
				},
			},
			&barbuilder.SpaceLarge{},
			makeDemo(update),
			&barbuilder.SpaceSmall{},
			makeCatalog(switcher, update),
		}
	})

	err := tb.Debug(config)
	if err != nil {
		panic(err)
	}

	err = tb.Uninstall()
	if err != nil {
		panic(err)
	}
}
