package types

import (
	"github.com/gorilla/websocket"
)

type Client interface {
	Connection() *websocket.Conn
	Manager() Manager
	Egress() chan Event

	GetClientID() string

	GetSession() string
	SetSession(string)

	PongHandler(string) error
	ReadMessages()
	WriteMessages()
}

type ClientRepo interface {
	GetAllClients() []Client
	AddClient(Client) error
	ClientExists(Client) (error, bool)
	DeleteClient(Client) error
	GetClient(string) Client
}
