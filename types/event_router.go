package types

type EventRouter interface {
	Route(event Event, client Client) error
	Handle(eventType EventType, eventHandler EventHandler) EventRouter
	Handlers() map[EventType]EventHandler
}
