package util

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func NewWebsocketUpgrader(rbs, wbs int, checkOrigin func(*http.Request) bool) *websocket.Upgrader {
	return &websocket.Upgrader{
		CheckOrigin:     checkOrigin,
		ReadBufferSize:  rbs,
		WriteBufferSize: wbs,
	}
}
