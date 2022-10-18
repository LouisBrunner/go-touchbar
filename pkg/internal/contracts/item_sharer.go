package contracts

type Sharer struct {
	CommonProperties

	ButtonImage string
	ButtonLabel string
	Enabled     bool

	// TODO: needs delegate or something?
}

func (me *Sharer) isAnItem() {}
