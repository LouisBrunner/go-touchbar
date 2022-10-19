package barbuilder

type Label struct {
	CommonProperties

	Content View

	// TODO: loads of options
}

func (me *Label) isAnItem() {}
