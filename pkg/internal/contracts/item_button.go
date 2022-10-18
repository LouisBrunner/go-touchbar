package contracts

type ButtonOnClick func()

type Button struct {
	CommonProperties
	View

	// TODO: finish

	OnClick ButtonOnClick
}

func (me *Button) isAnItem() {}
