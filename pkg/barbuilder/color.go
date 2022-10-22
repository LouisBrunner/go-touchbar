package barbuilder

type Color interface {
	isAColor()
}

type RGBAColor struct {
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

var _ Color = &RGBAColor{}

func (me *RGBAColor) isAColor() {}

type HexColor string

var _ Color = HexColor("")

func (me HexColor) isAColor() {}

// TODO: add standard colors https://developer.apple.com/documentation/appkit/nscolor/standard_colors?changes=_5&language=objc
// TODO: add UI colors https://developer.apple.com/documentation/appkit/nscolor/ui_element_colors?changes=_5&language=objc
// TODO: add more custom color spaces
