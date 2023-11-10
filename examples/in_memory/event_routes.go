package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/AdamShannag/gosockz/types"
)

// BroadcastMessageHandler will send out a message to all other participants in the session
func BroadcastMessageHandler(event types.Event, c types.Client) error {
	var chatevent BroadcastMessageEvent
	if err := json.Unmarshal(event.Payload, &chatevent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var broadMessage NewMessageEvent

	broadMessage.Sent = time.Now()
	broadMessage.Message = chatevent.Message
	broadMessage.From = chatevent.From

	data, err := json.Marshal(broadMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent types.Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage

	for _, client := range c.Manager().Clients().GetAllClients() {
		if client.GetSession() == c.GetSession() {
			client.Egress() <- outgoingEvent
		}

	}
	return nil
}

func SendMessageHandler(event types.Event, c types.Client) error {
	var msgEvent SendMessageEvent
	if err := json.Unmarshal(event.Payload, &msgEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	var msg NewMessageEvent

	msg.Sent = time.Now()
	msg.Message = msgEvent.Message
	msg.From = msgEvent.From

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal broadcast message: %v", err)
	}

	var outgoingEvent types.Event
	outgoingEvent.Payload = data
	outgoingEvent.Type = EventNewMessage

	reciever := c.Manager().Clients().GetClient(msgEvent.To)
	if reciever == nil {
		return fmt.Errorf("user: %v does not exist", msgEvent.To)
	}
	reciever.Egress() <- outgoingEvent

	return nil
}

func SessionHandler(event types.Event, c types.Client) error {
	var changeSessionEvent ChangeSessionEvent
	if err := json.Unmarshal(event.Payload, &changeSessionEvent); err != nil {
		return fmt.Errorf("bad payload in request: %v", err)
	}

	c.SetSession(changeSessionEvent.Name)

	return nil
}
