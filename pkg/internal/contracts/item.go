package contracts

type Items = []item

type item interface {
	isAnItem()
}

type ItemPriority string

const (
	ItemPriorityLow    ItemPriority = "low"
	ItemPriorityMedium ItemPriority = "medium"
	ItemPriorityHigh   ItemPriority = "high"
)

type CommonProperties struct {
	Visible   bool
	Priority  ItemPriority
	Principal bool
}
