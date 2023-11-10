package client

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func (c *ClientImpl) WriteMessages() {
	ticker := time.NewTicker(c.pongInterval)
	defer func() {
		ticker.Stop()
		c.manager.RemoveClient(c)
	}()

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}
			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Println(err)
			}

		case <-ticker.C:
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Println("writemsg: ", err)
				return
			}
		}

	}
}
