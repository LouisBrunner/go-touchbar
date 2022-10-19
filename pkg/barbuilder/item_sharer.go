package barbuilder

type Sharer struct {
	CommonProperties

	ButtonImage Image
	ButtonLabel string
	Disabled    bool

	// TODO: needs delegate or something?
}

func (me *Sharer) isAnItem() {}
