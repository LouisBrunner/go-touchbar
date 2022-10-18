package darwin

import (
	"unsafe"

	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
)

type touchBar struct {
	options  contracts.Options
	handlers *handlers
	context  unsafe.Pointer
}

func NewTouchBar(options contracts.Options) contracts.TouchBar {
	return &touchBar{
		options: options,
	}
}
