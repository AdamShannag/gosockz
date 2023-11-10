package types

import (
	"github.com/gorilla/websocket"
)

type Manager interface {
	Clients() ClientRepo
	EventRouter() EventRouter
	WebsocketUpgrader() *websocket.Upgrader

	AddClient(Client) error
	RemoveClient(Client)
}
