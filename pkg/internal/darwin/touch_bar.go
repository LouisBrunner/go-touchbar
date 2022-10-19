package darwin

import (
	"unsafe"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

type touchBar struct {
	options  barbuilder.Options
	handlers *handlers
	context  unsafe.Pointer
}

func NewTouchBar(options barbuilder.Options) barbuilder.TouchBar {
	return &touchBar{
		options: options,
	}
}
