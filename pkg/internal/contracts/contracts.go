package contracts

type TouchBar interface {
	Install() error

	// Debug is an alternative to `Install` which block the current thread to run a debug application
	Debug() error

	Update(config Configuration) error
	Uninstall() error
}

type Options struct {
	Configuration Configuration
}

type Configuration struct {
	Items  Items
	Escape *item
}
