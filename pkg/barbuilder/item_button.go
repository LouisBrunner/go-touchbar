package barbuilder

type ButtonOnClick func()

type Button struct {
	CommonProperties

	Title   string
	Image   Image
	OnClick ButtonOnClick

	// TODO: loads of options
}

func (me *Button) isAnItem() {}
