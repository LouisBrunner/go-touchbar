# go-touchbar [![Go Reference](https://pkg.go.dev/badge/github.com/LouisBrunner/go-touchbar.svg)](https://pkg.go.dev/github.com/LouisBrunner/go-touchbar)

Go library to integrate the MacBook Touch Bar

## Installation

```bash
go get github.com/LouisBrunner/go-touchbar
```

## Usage

```go
// Setup your window code (including NSApplication/NSWindow on macOS)

tb := touchbar.New(touchbar.Options{
  Configuration: touchbar.Configuration{
    // Add your configuration here
  },
})

err := tb.Install()
if err != nil {
  // handle
}

// run your application

err = tb.Uninstall()
if err != nil {
  // handle
}
```

TODO: add an example for the update flow

### Configuration

See [example application](./examples/tester/main.go) for a real-life example.

TODO: more details + godocs

## Further work

- Allow user customization
- Allow more options (check TODOs)
- Thread-safety
- Support custom images
- Better documentation
- Make a catalog like [Apple's](https://developer.apple.com/documentation/appkit/touch_bar/creating_and_customizing_the_touch_bar?language=objc)

## Acknowledgements

This library's API was influenced by [Electron's](https://www.electronjs.org/docs/latest/api/touch-bar).

The [example application](./examples/tester) is a reimplementation of [`electron-touch-bar`](https://github.com/pahund/electron-touch-bar).
