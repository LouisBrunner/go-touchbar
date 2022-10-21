package barbuilder

type SegmentedOnChange func(i int)

type Segment struct {
	Label string
	Image Image
}

type SegmentedControl struct {
	CommonProperties

	Segments       []Segment
	SelectMultiple bool
	OnChange       SegmentedOnChange
}

var _ Item = &SegmentedControl{}

func (me *SegmentedControl) isAnItem() {}
