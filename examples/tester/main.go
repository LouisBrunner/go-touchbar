package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"

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
	var onSpin func()

	reel1 := &barbuilder.Label{
		Content: &barbuilder.ContentLabel{},
	}
	reel2 := &barbuilder.Label{
		Content: &barbuilder.ContentLabel{},
	}
	reel3 := &barbuilder.Label{
		Content: &barbuilder.ContentLabel{},
	}
	result := &barbuilder.Label{
		Content: &barbuilder.ContentLabel{},
	}
	spin := &barbuilder.Button{
		Title:      "🎰 Spin",
		BezelColor: barbuilder.HexColor("#7851A9"),
		OnClick: func() {
			go onSpin()
		},
	}
	items := []barbuilder.Item{
		spin,
		&barbuilder.SpaceLarge{},
		reel1,
		&barbuilder.SpaceSmall{},
		reel2,
		&barbuilder.SpaceSmall{},
		reel3,
		&barbuilder.SpaceLarge{},
		result,
	}
	tb := touchbar.New(barbuilder.Options{})

	spinning := false
	mutex := sync.Mutex{}
	onSpin = func() {
		mutex.Lock()
		if spinning {
			mutex.Unlock()
			return
		}
		spinning = true
		mutex.Unlock()

		var err error
		defer func() {
			if err != nil {
				result.Content = &barbuilder.ContentLabel{
					Text:  fmt.Sprintf("Error: %v", err),
					Color: barbuilder.HexColor("#ff0000"),
				}
				tb.Update(barbuilder.Configuration{Items: items})
			}

			mutex.Lock()
			defer mutex.Unlock()
			spinning = false
		}()

		result.Content = &barbuilder.ContentLabel{}
		timeBetweenSpin := 10 * time.Millisecond
		start := time.Now()
		var value1, value2, value3 *string
		for {
			value1, err = getRandomValue()
			if err != nil {
				return
			}
			value2, err = getRandomValue()
			if err != nil {
				return
			}
			value3, err = getRandomValue()
			if err != nil {
				return
			}
			// FIXME: having to recreate the label is meh
			reel1.Content = &barbuilder.ContentLabel{Text: *value1}
			reel2.Content = &barbuilder.ContentLabel{Text: *value2}
			reel3.Content = &barbuilder.ContentLabel{Text: *value3}
			err = tb.Update(barbuilder.Configuration{Items: items})
			if err != nil {
				return
			}

			time.Sleep(timeBetweenSpin)
			if time.Since(start) >= 4*time.Second {
				break
			}
			timeBetweenSpin += timeBetweenSpin / 10
		}

		// FIXME: no easy way to get the last values...
		set := map[string]struct{}{
			*value1: {},
			*value2: {},
			*value3: {},
		}
		switch len(set) {
		case 1:
			result.Content = &barbuilder.ContentLabel{
				Text:  "💰 Jackpot!",
				Color: barbuilder.HexColor("#FDFF00"),
			}
		case 2:
			result.Content = &barbuilder.ContentLabel{
				Text:  "😍 Winner!",
				Color: barbuilder.HexColor("#FDFF00"),
			}
		default:
			result.Content = &barbuilder.ContentLabel{
				Text: "🙁 Spin Again",
			}
		}
		err = tb.Update(barbuilder.Configuration{Items: items})
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
