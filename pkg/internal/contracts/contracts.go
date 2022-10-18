package contracts

type TouchBar interface {
	Install() error

	// Debug is an alternative to `Install` which block the current thread to run a debug application
	Debug() error

	Update(config Configuration) error
	Uninstall() error
}

type Options struct {
	Configuration Configuration
}

type Configuration struct {
	Items  Items
	Escape *item
}

type Items = []item

type item interface {
	isAnItem()
}

type Button struct {
}

func (me *Button) isAnItem() {}

type ColorPicker struct {
}

func (me *ColorPicker) isAnItem() {}

type Group struct {
}

func (me *Group) isAnItem() {}

type Label struct {
}

func (me *Label) isAnItem() {}

type Popover struct {
}

func (me *Popover) isAnItem() {}

type Scrubber struct {
}

func (me *Scrubber) isAnItem() {}

type SegmentedControl struct {
}

func (me *SegmentedControl) isAnItem() {}

type Slider struct {
}

func (me *Slider) isAnItem() {}

type Spacer struct {
}

func (me *Spacer) isAnItem() {}
