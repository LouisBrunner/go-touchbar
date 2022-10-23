package main

import (
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
	"github.com/LouisBrunner/go-touchbar/pkg/barutils"
)

func makeCatalog(switcher barutils.Switcher, update func()) barbuilder.Item {
	// TODO: showcase Escape
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
			barutils.VirtualPopover(barbuilder.Popover{
				CollapsedText:  "Next",
				CollapsedImage: barbuilder.TBBookmarksTemplate,
				Bar: []barbuilder.Item{
					makePopoverCatalog(switcher),
					makeSliderCatalog(switcher, update),
				},
			}, switcher),
		},
	}, switcher)
}
