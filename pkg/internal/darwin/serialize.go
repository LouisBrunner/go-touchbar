package darwin

import (
	"fmt"

	"github.com/LouisBrunner/go-touchbar/pkg/barbuilder"
)

type item interface{}

type itemGroup struct {
	barbuilder.CommonProperties

	Direction          barbuilder.GroupDirection
	Children           []string
	PrefersEqualWidth  bool
	PreferredItemWidth float32
}

type itemPopover struct {
	barbuilder.CommonProperties

	Collapsed       barbuilder.View
	Bar             []string
	ShowCloseButton bool
	PressAndHold    bool
}

type flatConfig struct {
	Principal       string
	Default         []string
	Items           map[string]item
	OtherItemsProxy bool
	Escape          string
}

const namespace = "net.lbrunner.touchbar"

const (
	standardSpaceSmall      = namespace + ".small_space"
	standardSpaceLarge      = namespace + ".large_space"
	standardSpaceFlexible   = namespace + ".flexible_space"
	standardCandidateList   = namespace + ".candidates"
	standardCharacterPicker = namespace + ".char_picker"
	standardTextFormat      = namespace + ".text_format"
	standardTextAlignment   = namespace + ".text_align"
	standardTextColorPicker = namespace + ".text_color"
	standardTextList        = namespace + ".text_list"
	standardTextStyle       = namespace + ".text_style"
)

func makeID(prefix, kind string, i int) string {
	return fmt.Sprintf("%s.%s.%s.%d", namespace, prefix, kind, i)
}

func processItem(prefix string, i int, item barbuilder.Item, principal *string, dict map[string]item, handlers *handlers) (string, interface{}, error) {
	id := ""
	isPrincipal := false
	var result interface{} = item

	switch widget := item.(type) {
	// customizable
	case *barbuilder.Button:
		id = makeID(prefix, "button", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.buttons[id] = widget.OnClick

	case *barbuilder.Candidates:
		id = makeID(prefix, "candidates", i)
		isPrincipal = widget.Principal

	case *barbuilder.ColorPicker:
		id = makeID(prefix, "colorpicker", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.colorPickers[id] = widget.OnSelected

	case *barbuilder.Custom:
		id = makeID(prefix, "custom", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.customs[id] = widget.OnEvent

	case *barbuilder.Group:
		id = makeID(prefix, "group", i)
		isPrincipal = widget.Principal
		list, principal, err := processItems(fmt.Sprintf("%s.group.%d", prefix, i), widget.Children, dict, handlers)
		if err != nil {
			return "", nil, err
		}
		if principal != "" {
			return "", nil, fmt.Errorf("principal is not supported in sub touch bars")
		}
		result = itemGroup{
			CommonProperties:   widget.CommonProperties,
			Direction:          widget.Direction,
			Children:           list,
			PrefersEqualWidth:  widget.PrefersEqualWidth,
			PreferredItemWidth: widget.PreferredItemWidth,
		}

	case *barbuilder.Label:
		id = makeID(prefix, "label", i)
		isPrincipal = widget.Principal

	case *barbuilder.Picker:
		id = makeID(prefix, "picker", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.pickers[id] = widget.OnSelected

	case *barbuilder.Popover:
		id = makeID(prefix, "popover", i)
		isPrincipal = widget.Principal
		list, principal, err := processItems(fmt.Sprintf("%s.popover.%d", prefix, i), widget.Bar, dict, handlers)
		if err != nil {
			return "", nil, err
		}
		if principal != "" {
			return "", nil, fmt.Errorf("principal is not supported in sub touch bars")
		}
		result = itemPopover{
			CommonProperties: widget.CommonProperties,
			Collapsed:        widget.Collapsed,
			Bar:              list,
			ShowCloseButton:  widget.ShowCloseButton,
			PressAndHold:     widget.PressAndHold,
		}

	case *barbuilder.Scrubber:
		id = makeID(prefix, "scrubber", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.scrubbers[id] = widget.OnChange

	case *barbuilder.SegmentedControl:
		id = makeID(prefix, "segmented", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.segments[id] = widget.OnChange

	case *barbuilder.Sharer:
		id = makeID(prefix, "sharer", i)
		isPrincipal = widget.Principal

	case *barbuilder.Slider:
		id = makeID(prefix, "slider", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.sliders[id] = widget.OnChange

	case *barbuilder.Stepper:
		id = makeID(prefix, "stepper", i)
		isPrincipal = widget.Principal
		if handlers == nil {
			return "", nil, fmt.Errorf("cannot use this item in this context %T (%v)", item, item)
		}
		handlers.steppers[id] = widget.OnChange

		// standards
	case *barbuilder.SpaceSmall:
		id = standardSpaceSmall
	case *barbuilder.SpaceLarge:
		id = standardSpaceLarge
	case *barbuilder.SpaceFlexible:
		id = standardSpaceFlexible
	case *barbuilder.CharacterPicker:
		id = standardCharacterPicker
	case *barbuilder.CandidateList:
		id = standardCandidateList
	case *barbuilder.TextFormat:
		id = standardTextFormat
	case *barbuilder.TextAlignment:
		id = standardTextAlignment
	case *barbuilder.TextColorPicker:
		id = standardTextColorPicker
	case *barbuilder.TextList:
		id = standardTextList
	case *barbuilder.TextStyle:
		id = standardTextStyle

	default:
		return "", nil, fmt.Errorf("unknown item type %T (%v)", item, item)
	}

	if isPrincipal {
		if *principal != "" {
			return "", nil, fmt.Errorf("duplicate principal: %v vs %v", *principal, id)
		}
		*principal = id
	}

	return id, result, nil
}

func processItems(prefix string, items []barbuilder.Item, dict map[string]item, handlers *handlers) ([]string, string, error) {
	principal := ""
	list := make([]string, 0, len(items))

	for i, item := range items {
		id, result, err := processItem(prefix, i, item, &principal, dict, handlers)
		if err != nil {
			return nil, "", err
		}

		if id == "" {
			return nil, "", fmt.Errorf("id not generated properly for %T (%v)", item, item)
		}

		dict[id] = result
		list = append(list, id)
	}

	return list, principal, nil
}

func processConfig(config *barbuilder.Configuration) (*flatConfig, *handlers, error) {
	dict := make(map[string]item, len(config.Items))
	handlers := handlers{
		buttons:      make(map[string]barbuilder.ButtonOnClick),
		colorPickers: make(map[string]barbuilder.ColorPickerOnSelected),
		customs:      make(map[string]barbuilder.CustomOnEvent),
		pickers:      make(map[string]barbuilder.PickerOnSelected),
		scrubbers:    make(map[string]barbuilder.ScrubberOnChange),
		segments:     make(map[string]barbuilder.SegmentedOnChange),
		sliders:      make(map[string]barbuilder.SliderOnChange),
		steppers:     make(map[string]barbuilder.StepperOnChange),
	}

	list, principal, err := processItems("", config.Items, dict, &handlers)
	if err != nil {
		return nil, nil, err
	}

	escapeID := ""
	if config.Escape != nil {
		id, result, err := processItem("escape", 0, *config.Escape, &principal, dict, nil)
		if err != nil {
			return nil, nil, err
		}
		dict[id] = result
		escapeID = id
	}

	data := flatConfig{
		Principal: principal,
		Default:   list,
		Items:     dict,
		// FIXME: should we allow positioning for that?
		OtherItemsProxy: config.OtherItemsProxy,
		Escape:          escapeID,
	}
	return &data, &handlers, nil
}
