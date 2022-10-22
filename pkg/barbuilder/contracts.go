package barbuilder

type TouchBar interface {
	Install(config Configuration) error

	// Debug is an alternative to `Install` which block the current thread to run a debug application
	Debug(config Configuration) error

	Update(config Configuration) error
	Uninstall() error
}

type Options struct {
}

type Configuration struct {
	Items           []Item
	Escape          *Item
	OtherItemsProxy bool
}
