package router

import (
	"errors"

	"github.com/AdamShannag/gosockz/types"
)

var (
	ErrEventNotSupported = errors.New("this event type is not supported")
)

type EventRouter struct {
	handlers map[types.EventType]types.EventHandler
}

func NewEventRouter() *EventRouter {
	return &EventRouter{handlers: make(map[types.EventType]types.EventHandler)}
}

func (er *EventRouter) Route(event types.Event, client types.Client) error {
	if handler, ok := er.handlers[event.Type]; ok {
		if err := handler(event, client); err != nil {
			return err
		}
		return nil
	} else {
		return ErrEventNotSupported
	}
}

func (er *EventRouter) Handle(eventType types.EventType, eventHandler types.EventHandler) types.EventRouter {
	er.handlers[eventType] = eventHandler
	return er
}

func (er *EventRouter) Handlers() map[types.EventType]types.EventHandler {
	return er.handlers
}
