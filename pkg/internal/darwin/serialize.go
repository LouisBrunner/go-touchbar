package darwin

import (
	"fmt"

	"github.com/LouisBrunner/go-touchbar/pkg/internal/contracts"
)

type item interface{}

type flatConfig struct {
	Principal string
	Default   []string
	Items     map[string]item
	Escape    item
}

const namespace = "net.lbrunner.touchbar"

func makeID(kind string, i int) string {
	return fmt.Sprintf("%s.%s.%d", namespace, kind, i)
}

func processConfig(config *contracts.Configuration) (*flatConfig, *handlers, error) {
	list := make([]string, 0, len(config.Items))
	dict := make(map[string]item, len(config.Items))
	handlers := handlers{
		buttons:      make(map[string]contracts.ButtonOnClick),
		colorPickers: make(map[string]contracts.ColorPickerOnSelected),
		customs:      make(map[string]contracts.CustomOnEvent),
		pickers:      make(map[string]contracts.PickerOnSelected),
	}
	principal := ""

	for i, item := range config.Items {
		id := ""
		isPrincipal := false

		switch widget := item.(type) {
		case *contracts.Button:
			id = makeID("button", i)
			isPrincipal = widget.Principal
			handlers.buttons[id] = widget.OnClick
		case *contracts.Candidates:
			id = makeID("candidates", i)
			isPrincipal = widget.Principal
		case *contracts.ColorPicker:
			id = makeID("colorpicker", i)
			isPrincipal = widget.Principal
			handlers.colorPickers[id] = widget.OnSelected
		case *contracts.Custom:
			id = makeID("custom", i)
			isPrincipal = widget.Principal
			handlers.customs[id] = widget.OnEvent
		case *contracts.Group:
			id = makeID("group", i)
			isPrincipal = widget.Principal
		case *contracts.Label:
			id = makeID("label", i)
			isPrincipal = widget.Principal
		case *contracts.Picker:
			id = makeID("picker", i)
			isPrincipal = widget.Principal
			handlers.pickers[id] = widget.OnSelected
		case *contracts.Popover:
			id = makeID("popover", i)
			isPrincipal = widget.Principal
		case *contracts.Scrubber:
			id = makeID("scrubber", i)
			isPrincipal = widget.Principal
		case *contracts.SegmentedControl:
			id = makeID("segmented", i)
			isPrincipal = widget.Principal
		case *contracts.Sharer:
			id = makeID("sharer", i)
			isPrincipal = widget.Principal
		case *contracts.Slider:
			id = makeID("slider", i)
			isPrincipal = widget.Principal
		case *contracts.Spacer:
			id = makeID("spacer", i)
			isPrincipal = widget.Principal
		case *contracts.Stepper:
			id = makeID("stepper", i)
			isPrincipal = widget.Principal
		default:
			return nil, nil, fmt.Errorf("unknown item type %T (%v)", item, item)
		}

		if id == "" {
			return nil, nil, fmt.Errorf("id not generated properly for %T (%v)", item, item)
		}
		if isPrincipal {
			if principal != "" {
				return nil, nil, fmt.Errorf("duplicate principal: %v vs %v", principal, id)
			}
			principal = id
		}

		dict[id] = item
		list = append(list, id)
	}

	data := flatConfig{
		Principal: principal,
		Default:   list,
		Items:     dict,
		Escape:    config.Escape,
	}
	return &data, &handlers, nil
}
