# go-touchbar

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

### Configuration

TODO: finish

## Further work

- Allow user customization
- Allow more options (check TODOs)
- Thread-safety
- Support custom images

## Acknowledgements

This library's API was influenced by [Electron's](https://www.electronjs.org/docs/latest/api/touch-bar).

The [testing application](./examples/tester) is a reimplementation of [`electron-touch-bar`](https://github.com/pahund/electron-touch-bar).
