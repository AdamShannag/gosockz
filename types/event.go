package types

import "encoding/json"

type Event struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(Event, Client) error

type EventType string
