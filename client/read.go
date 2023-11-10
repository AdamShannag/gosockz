package client

import (
	"encoding/json"
	"log"
	"time"

	"github.com/AdamShannag/gosockz/types"
	"github.com/gorilla/websocket"
)

func (c *ClientImpl) ReadMessages() {
	defer func() {
		c.manager.RemoveClient(c)
	}()

	c.connection.SetReadLimit(c.readLimit)

	if err := c.connection.SetReadDeadline(time.Now().Add(c.pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.connection.SetPongHandler(c.PongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message: %v", err)
			}
			break
		}

		var request types.Event
		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling message: %v", err)
			break
		}

		if err := c.manager.EventRouter().Route(request, c); err != nil {
			log.Println("error handeling Message: ", err)
		}
	}
}
