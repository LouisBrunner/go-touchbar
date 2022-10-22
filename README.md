# go-touchbar [![Go Reference](https://pkg.go.dev/badge/github.com/LouisBrunner/go-touchbar.svg)](https://pkg.go.dev/github.com/LouisBrunner/go-touchbar)

Go library to integrate the MacBook Touch Bar

## Installation

```bash
go get github.com/LouisBrunner/go-touchbar
```

## Usage

```go
// Setup your window code (including NSApplication/NSWindow on macOS)

tb := touchbar.New(touchbar.Options{})

err := tb.Install(Configuration: touchbar.Configuration{
  // Add your configuration here
})
if err != nil {
  // handle
}

// run your application

// when you want to update the touchbar (even from another routine), call do
err = tb.Update(Configuration: touchbar.Configuration{
  // Add your updated configuration here
})
if err != nil {
  // handle
}

err = tb.Uninstall()
if err != nil {
  // handle
}
```

### Configuration

See [example application](./examples/tester/main.go) for a real-life example.

TODO: more details + godocs

## Further work

Check TODO/FIXME as well

- (!!!) Finish implementing widgets
- Allow user customization (`customizationLabel`, `templateItems`, etc)
- Ability to use standard/UI colors
- Better choices for the color-picker
- More customization to the custom widget
- Group compressions
- Layout constraints (e.g. sizing)
- Support custom images
- Thread-safety
- Better documentation
- Better validation in Go (validator on the structs?)
- Make a catalog like [Apple's](https://developer.apple.com/documentation/appkit/touch_bar/creating_and_customizing_the_touch_bar?language=objc)

## Acknowledgements

This library's API was influenced by [Electron's](https://www.electronjs.org/docs/latest/api/touch-bar).

The [example application](./examples/tester) is a reimplementation of [`electron-touch-bar`](https://github.com/pahund/electron-touch-bar).
