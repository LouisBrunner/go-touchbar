package contracts

type PickerOnSelected func(i int)

type Picker struct {
	CommonProperties

	Items           []View
	SingleSelection bool
	Collapsed       bool
	OnSelected      PickerOnSelected
}

func (me *Picker) isAnItem() {}
