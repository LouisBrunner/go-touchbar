package barbuilder

type View interface {
	isAView()
}

type ContentLabel struct {
	Text  string
	Color Color
}

var _ View = &ContentLabel{}

func (me *ContentLabel) isAView() {}

type ContentImage struct {
	Image Image
}

var _ View = &ContentImage{}

func (me *ContentImage) isAView() {}
