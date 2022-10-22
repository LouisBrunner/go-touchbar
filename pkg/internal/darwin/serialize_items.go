package darwin

import (
	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

type item interface{}

type itemButton struct {
	barbuilder.CommonProperties

	Title      string
	Image      barbuilder.Image
	Disabled   bool
	BezelColor barbuilder.Color
}

type itemGroup struct {
	barbuilder.CommonProperties

	Direction          barbuilder.GroupDirection
	Children           []identifier
	PrefersEqualWidth  bool
	PreferredItemWidth float32
}

type itemPopover struct {
	barbuilder.CommonProperties

	CollapsedText  string
	CollapsedImage barbuilder.Image
	Bar            []identifier
	PressAndHold   bool
}
