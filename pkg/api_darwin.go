package touchbar

import (
	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
	"github.com/LouisBrunner/go-touchbar/pkg/internal/darwin"
)

type TouchBar = contracts.TouchBar
type Options = contracts.Options
type Configuration = contracts.Configuration
type Items = contracts.Items
type Button = contracts.Button
type ColorPicker = contracts.ColorPicker
type Group = contracts.Group
type Label = contracts.Label
type Popover = contracts.Popover
type Scrubber = contracts.Scrubber
type SegmentedControl = contracts.SegmentedControl
type Slider = contracts.Slider
type Spacer = contracts.Spacer

var New = darwin.NewTouchBar
