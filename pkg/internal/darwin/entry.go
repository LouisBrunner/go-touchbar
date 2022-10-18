package darwin

import (
	"encoding/json"
	"fmt"
	"unsafe"

	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
)

//#cgo CFLAGS: -x objective-c -std=c2x
//#cgo LDFLAGS: -framework Foundation -framework Cocoa
//#include "entry.h"
import "C"

func serializeConfig(config *contracts.Configuration) (*C.char, *handlers, error) {
	data, handlers, err := processConfig(config)
	if err != nil {
		return nil, nil, err
	}
	buffer, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}
	return C.CString(string(buffer)), handlers, nil
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

//export handleEvent
func handleEvent(raw unsafe.Pointer, event *C.char) {
	me := (*touchBar)(raw)
	me.handleEvent(C.GoString(event))
}

func (me *touchBar) install(debug bool) error {
	if me.context != nil {
		return fmt.Errorf("touch bar already initialized")
	}
	mode := C.kMainWindow
	if debug {
		mode = C.kDebug
	}
	data, handlers, err := serializeConfig(&me.options.Configuration)
	if err != nil {
		return err
	}
	defer C.free(unsafe.Pointer(data))
	me.handlers = handlers
	result := C.initTouchBar(C.AttachMode(mode), data, unsafe.Pointer(me))
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
	data, handlers, err := serializeConfig(&me.options.Configuration)
	if err != nil {
		return err
	}
	defer C.free(unsafe.Pointer(data))
	me.handlers = handlers
	return handleError(C.updateTouchBar(me.context, data))
}

func (me *touchBar) Uninstall() error {
	if me.context == nil {
		return fmt.Errorf("touch bar has not been initialized")
	}
	return handleError(C.destroyTouchBar(me.context))
}
