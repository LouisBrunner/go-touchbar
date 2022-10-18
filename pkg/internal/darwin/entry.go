package darwin

import (
	"fmt"
	"unsafe"

	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
)

//#cgo CFLAGS: -x objective-c -std=c2x
//#cgo LDFLAGS: -framework Foundation -framework Cocoa
//#include "entry.h"
import "C"

type touchBar struct {
	options contracts.Options
	context unsafe.Pointer
}

func serializeConfig(config *contracts.Configuration) C.TouchBar {
	// TODO: how?
	return C.TouchBar{}
}

func transformError(err *C.char) error {
	return fmt.Errorf("darwin glue: %v", C.GoString(err))
}

func handleError(result C.ErrorResult) error {
	if result.err != nil {
		return transformError(result.err)
	}
	return nil
}

func (me *touchBar) install(debug bool) error {
	if me.context != nil {
		return fmt.Errorf("touch bar already initialized")
	}
	mode := C.kMainWindow
	if debug {
		mode = C.kDebug
	}
	result := C.initTouchBar(C.AttachMode(mode), serializeConfig(&me.options.Configuration))
	if result.err != nil {
		return transformError(result.err)
	}
	me.context = result.result
	return nil
}

func (me *touchBar) Install() error {
	return me.install(false)
}

func (me *touchBar) Debug() error {
	err := me.install(true)
	if err != nil {
		return err
	}
	return handleError(C.runDebug(me.context))
}

func (me *touchBar) Update(configuration contracts.Configuration) error {
	if me.context == nil {
		return fmt.Errorf("touch bar has not been initialized")
	}
	return handleError(C.updateTouchBar(me.context, serializeConfig(&configuration)))
}

func (me *touchBar) Uninstall() error {
	if me.context == nil {
		return fmt.Errorf("touch bar has not been initialized")
	}
	return handleError(C.destroyTouchBar(me.context))
}

func NewTouchBar(options contracts.Options) contracts.TouchBar {
	return &touchBar{
		options: options,
	}
}
