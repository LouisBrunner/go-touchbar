package barbuilder

type Item interface {
	isAnItem()
}

type ItemPriority float32

const (
	ItemPriorityLow    ItemPriority = -1000
	ItemPriorityMedium ItemPriority = 0
	ItemPriorityHigh   ItemPriority = 1000
)

type CommonProperties struct {
	Visible   bool
	Priority  ItemPriority
	Principal bool
}
