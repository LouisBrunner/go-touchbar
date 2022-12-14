package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

func makeDemo(update func()) barbuilder.Item {
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
				update()
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
			update()

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
		update()
	}

	return &barbuilder.Popover{
		CollapsedText:  "Demo",
		CollapsedImage: barbuilder.SFSymbol("dollarsign.circle"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Demo",
				},
			},
			&barbuilder.SpaceLarge{},
			spin,
			&barbuilder.SpaceLarge{},
			reel1,
			&barbuilder.SpaceSmall{},
			reel2,
			&barbuilder.SpaceSmall{},
			reel3,
			&barbuilder.SpaceLarge{},
			result,
		},
	}
}
