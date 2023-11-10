package main

import "time"

const (
	// EventBroadcastMessage is the event name for new chat messages to be broadcasted to all other clients
	EventBroadcastMessage = "broadcast_message"
	// EventSendMessage is the event name for new chat messages sent
	EventSendMessage = "send_message"
	// EventNewMessage is a response to send_message
	EventNewMessage = "new_message"
	// EventChangeSession is event when switching rooms
	EventChangeSession = "change_session"
)

type BroadcastMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
}

type NewMessageEvent struct {
	Message string    `json:"message"`
	From    string    `json:"from"`
	Sent    time.Time `json:"sent"`
}

type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"from"`
	To      string `json:"to"`
}

type ChangeSessionEvent struct {
	Name string `json:"name"`
}
