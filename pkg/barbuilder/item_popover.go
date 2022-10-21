package barbuilder

type Popover struct {
	CommonProperties

	Collapsed       View
	Bar             []Item
	ShowCloseButton bool
	PressAndHold    bool
}

var _ Item = &Popover{}

func (me *Popover) isAnItem() {}
