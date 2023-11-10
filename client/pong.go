package client

import (
	"time"
)

func (c *ClientImpl) PongHandler(pongMsg string) error {
	return c.connection.SetReadDeadline(time.Now().Add(c.pongWait))
}
