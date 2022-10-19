package barbuilder

type View interface {
	isAView()
}

type ContentLabel struct {
	Text string
}

func (me *ContentLabel) isAView() {}

type ContentImage struct {
	Image Image
}

func (me *ContentImage) isAView() {}
