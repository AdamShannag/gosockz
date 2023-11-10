package client

import (
	"time"

	"github.com/AdamShannag/gosockz/types"
	"github.com/gorilla/websocket"
)

type ClientImpl struct {
	clientID               string
	connection             *websocket.Conn
	manager                types.Manager
	egress                 chan types.Event
	chatroom               string
	pongWait, pongInterval time.Duration
	readLimit              int64
}

func (c *ClientImpl) Connection() *websocket.Conn {
	return c.connection
}
func (c *ClientImpl) Manager() types.Manager {
	return c.manager
}
func (c *ClientImpl) GetSession() string {
	return c.chatroom
}
func (c *ClientImpl) SetSession(room string) {
	c.chatroom = room
}
func (c *ClientImpl) Egress() chan types.Event {
	return c.egress
}

func (c *ClientImpl) GetClientID() string {
	return c.clientID
}

func NewClient(clientID string, conn *websocket.Conn, manager types.Manager, pongWait, pongInterval time.Duration) types.Client {
	return &ClientImpl{
		clientID:     clientID,
		connection:   conn,
		manager:      manager,
		egress:       make(chan types.Event),
		pongWait:     pongWait,
		pongInterval: pongInterval,
	}
}

func NewDefaultClient(clientID string, conn *websocket.Conn, manager types.Manager) types.Client {
	return &ClientImpl{
		clientID:     clientID,
		connection:   conn,
		manager:      manager,
		egress:       make(chan types.Event),
		pongWait:     10 * time.Second,
		pongInterval: 9 * time.Second,
		readLimit:    512,
	}
}
