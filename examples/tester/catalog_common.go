package main

import (
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
	"github.com/LouisBrunner/go-touchbar/pkg/barutils"
)

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
