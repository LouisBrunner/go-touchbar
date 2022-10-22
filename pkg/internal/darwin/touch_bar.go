package darwin

import (
	"sync"
	"unsafe"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

type touchBar struct {
	options  barbuilder.Options
	handlers *handlers
	context  unsafe.Pointer
	lock     sync.Mutex
}

func NewTouchBar(options barbuilder.Options) barbuilder.TouchBar {
	return &touchBar{
		options: options,
	}
}
