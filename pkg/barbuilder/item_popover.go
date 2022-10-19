package barbuilder

type Popover struct {
	CommonProperties

	Collapsed       View
	Bar             []Item
	ShowCloseButton bool
	PressAndHold    bool
}

func (me *Popover) isAnItem() {}
