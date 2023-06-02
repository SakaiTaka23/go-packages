package main

import (
	"wire/wire"
)

func main() {
	eventInst := wire.InitializeEvent()
	eventInst.Start()

	customEventInst := wire.InitializeCustomEvent("Hello world!")
	customEventInst.Start()
}
