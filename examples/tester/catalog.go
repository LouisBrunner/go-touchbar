package main

import (
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
	"github.com/LouisBrunner/go-touchbar/pkg/barutils"
)

func makeCatalog(switcher barutils.Switcher, update func()) barbuilder.Item {
	return barutils.VirtualPopover(barbuilder.Popover{
		CollapsedText:  "Catalog",
		CollapsedImage: barbuilder.TBBookmarksTemplate,
		Bar: []barbuilder.Item{
			&barbuilder.Label{
				Content: &barbuilder.ContentLabel{
					Text: "Catalog",
				},
			},
			makeLabelCatalog(),
			makeButtonCatalog(update),
			makePopoverCatalog(switcher),
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
