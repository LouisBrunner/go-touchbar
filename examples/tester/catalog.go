package main

import (
	"fmt"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
	"github.com/LouisBrunner/go-touchbar/pkg/barutils"
)

func makeCatalog(switcher barutils.Switcher, update func()) barbuilder.Item {
	// TODO: showcase Escape
	// TODO: showcase standards
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Catalog",
		CollapsedImage: barbuilder.TBBookmarksTemplate,
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Catalog",
				},
			},
			makeCommonCatalog(switcher),
			makeLabelCatalog(),
			makeButtonCatalog(update),
			makePopoverCatalog(switcher),
			makeSliderCatalog(switcher, update),
		},
	}, switcher)
}

func makeCommonCatalog(switcher barutils.Switcher) barbuilder.Item {
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Common",
		CollapsedImage: barbuilder.SFSymbol("list.star"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Common",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Principal",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Before",
						},
					},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Principal: true,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Principal",
						},
					},
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "After",
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Priority",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very High",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityLow * 2,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Very Low",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityMedium,
						},
						Content: &barbuilder.ContentLabel{
							Text: "Medium",
						},
					},
					&barbuilder.SpaceLarge{},
					&barbuilder.Label{
						CommonProperties: barbuilder.CommonProperties{
							Priority: barbuilder.ItemPriorityHigh,
						},
						Content: &barbuilder.ContentLabel{
							Text: "VERY LONG TEXT WHICH WILL HIDE THE OTHERS",
						},
					},
				},
			},
		},
	}, switcher)
}

func makeLabelCatalog() barbuilder.Item {
	return &barbuilder.Popover{
		CollapsedText:  "Label",
		CollapsedImage: barbuilder.SFSymbol("text.alignleft"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Label",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text:  "Color #f42309",
					Color: barbuilder.HexColor("#f42309"),
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Touch Bar Icon:",
				},
			},
			&barbuilder.Label{
				Content: &barbuilder.ContentImage{
					Image: barbuilder.TBAddTemplate,
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "SF Symbol:",
				},
			},
			&barbuilder.Label{
				Content: &barbuilder.ContentImage{
					Image: barbuilder.SFSymbol("text.bubble"),
				},
			},
		},
	}
}

func makeButtonCatalog(update func()) barbuilder.Item {
	result := &barbuilder.Label{Content: &barbuilder.ContentLabel{Text: ""}}

	return &barbuilder.Popover{
		CollapsedText:  "Button",
		CollapsedImage: barbuilder.SFSymbol("hand.point.up.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Button",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title: "Plain",
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button1 pressed"}
					update()
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Image: barbuilder.TBAlarmTemplate,
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button2 pressed"}
					update()
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title:      "With Icon & Color",
				Image:      barbuilder.SFSymbol("exclamationmark.triangle.fill"),
				BezelColor: barbuilder.HexColor("#e35412"),
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button3 pressed"}
					update()
				},
			},
			&barbuilder.SpaceSmall{},
			result,
			&barbuilder.SpaceSmall{},
			&barbuilder.Button{
				Title:    "Disabled",
				Image:    barbuilder.SFSymbol("sunrise.fill"),
				Disabled: true,
				OnClick: func() {
					result.Content = &barbuilder.ContentLabel{Text: "Button4 pressed"}
					update()
				},
			},
		},
	}
}

func makePopoverCatalog(switcher barutils.Switcher) barbuilder.Item {
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Popover",
		CollapsedImage: barbuilder.SFSymbol("bubble.left.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Popover",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Click Once",
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Releasing doesn't dismiss, click the X when you are done",
						},
					},
					&barbuilder.Button{
						Title: "Click Me",
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedImage: barbuilder.SFSymbol("escape"),
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Releasing doesn't dismiss, click the X when you are done",
						},
					},
					&barbuilder.Button{
						Title: "Click Me",
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText:  "Press & Hold",
				CollapsedImage: barbuilder.SFSymbol("rectangle.compress.vertical"),
				PressAndHold:   true,
				Bar: []barbuilder.Item{
					&barbuilder.Label{
						Content: &barbuilder.ContentLabel{
							Text: "Keep pressing!",
						},
					},
					&barbuilder.Button{
						Title: "Click me by sliding and releasing over me",
					},
				},
			},
		},
	}, switcher)
}

func makeSliderCatalog(switcher barutils.Switcher, update func()) barbuilder.Item {
	result := &barbuilder.Label{Content: &barbuilder.ContentLabel{Text: ""}}

	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Slider",
		CollapsedImage: barbuilder.SFSymbol("ruler.fill"),
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Slider",
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "Simple",
				PressAndHold:  true,
				Bar: []barbuilder.Item{
					&barbuilder.Slider{
						MinimumValue: 67,
						StartValue:   90,
						MaximumValue: 100,
						OnChange: func(value float64) {
							result.Content = &barbuilder.ContentLabel{Text: fmt.Sprintf("value: %v", value)}
							update()
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			&barbuilder.Popover{
				CollapsedText: "All options",
				PressAndHold:  true,
				Bar: []barbuilder.Item{
					&barbuilder.Slider{
						Label:            "Normal",
						MinimumValue:     0,
						StartValue:       7,
						MaximumValue:     10,
						MinimumAccessory: barbuilder.TBAlarmTemplate,
						MaximumAccessory: barbuilder.TBAddDetailTemplate,
						AccessoryWidth:   barbuilder.SliderAccessoryWide,
						OnChange: func(value float64) {
							result.Content = &barbuilder.ContentLabel{Text: fmt.Sprintf("value: %v", value)}
							update()
						},
					},
				},
			},
			&barbuilder.SpaceSmall{},
			result,
		},
	}, switcher)
}
