//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"wire/event"
	"wire/greeter"
	"wire/message"
)

func InitializeEvent() event.Event {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewMessage)
	return event.Event{}
}

func InitializeCustomEvent(phrase string) event.Event {
	wire.Build(event.NewEvent, greeter.NewGreeter, message.NewCustomMessage)
	return event.Event{}
}
