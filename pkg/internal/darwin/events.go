package darwin

import (
	"encoding/json"
	"fmt"

	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
)

type handlers struct {
	buttons      map[string]contracts.ButtonOnClick
	colorPickers map[string]contracts.ColorPickerOnSelected
	customs      map[string]contracts.CustomOnEvent
	pickers      map[string]contracts.PickerOnSelected
}

const (
	eventButton      = "button"
	eventColorPicker = "color_picker"
	eventCustom      = "custom"
	eventPicker      = "picker"
)

type event struct {
	Kind   string
	Target string
	Data   json.RawMessage
}

func (me *touchBar) handleEventLogic(eventJSON string) error {
	event := event{}
	err := json.Unmarshal([]byte(eventJSON), &event)
	if err != nil {
		return err
	}

	switch event.Kind {
	case eventButton:
		handler, found := me.handlers.buttons[event.Target]
		if !found {
			return fmt.Errorf("unknown button %v", event.Target)
		}
		handler()

	case eventColorPicker:
		handler, found := me.handlers.colorPickers[event.Target]
		if !found {
			return fmt.Errorf("unknown color picker %v", event.Target)
		}
		data := contracts.ColorPickerColor{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventCustom:
		handler, found := me.handlers.customs[event.Target]
		if !found {
			return fmt.Errorf("unknown custom %v", event.Target)
		}
		data := contracts.CustomEvent{}
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	case eventPicker:
		handler, found := me.handlers.pickers[event.Target]
		if !found {
			return fmt.Errorf("unknown picker %v", event.Target)
		}
		data := 0
		err := json.Unmarshal(event.Data, &data)
		if err != nil {
			return err
		}
		handler(data)

	default:
		return fmt.Errorf("unknown kind %v", event.Kind)
	}

	return nil
}

func (me *touchBar) handleEvent(eventJSON string) {
	err := me.handleEventLogic(eventJSON)
	if err != nil {
		// TODO: no idea what to do, needs some kind of logger?
		fmt.Printf("event error: %v\n", err)
		return
	}
}
