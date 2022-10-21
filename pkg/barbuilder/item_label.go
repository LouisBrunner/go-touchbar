package barbuilder

type Label struct {
	CommonProperties

	Content View

	// TODO: loads of options
}

var _ Item = &Label{}

func (me *Label) isAnItem() {}
