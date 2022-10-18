package contracts

type Popover struct {
	CommonProperties

	Collapsed       View
	Bar             Items
	ShowCloseButton bool
	PressAndHold    bool
}

func (me *Popover) isAnItem() {}
