package contracts

type ColorPickerKind string

const (
	ColorPickerStandard ColorPickerKind = "standard"
	ColorPickerText     ColorPickerKind = "text"
	ColorPickerStroke   ColorPickerKind = "stroke"
)

type ColorPickerColor struct {
	RGB   string
	Alpha float32
}

type ColorPickerOnSelected func(color ColorPickerColor)

type ColorPicker struct {
	CommonProperties

	Kind       ColorPickerKind
	ShowsAlpha bool
	Enabled    bool
	OnSelected ColorPickerOnSelected

	// TODO: custom color list
	// TODO: custom color spaces
}

func (me *ColorPicker) isAnItem() {}
