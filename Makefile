all:
.PHONY: all

build-example:
	go build -o examples/Test.app/Contents/MacOS/tester ./examples/tester
.PHONY: build-example

run-example: build-example
	open -W examples/Test.app
.PHONY: run-example