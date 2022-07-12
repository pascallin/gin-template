package events

import (
	"github.com/teris-io/shortid"
	"time"
)

// EventPayload is a generic key/value map for sending out to chat clients.
type EventPayload map[string]interface{}

// Event is any kind of event.  A type is required to be specified.
type Event struct {
	Type      EventType `json:"type,omitempty"`
	ID        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
	Event     EventType `json:"event"`
}

// SetDefaults will set default properties of all inbound events.
func (e *Event) SetDefaults() {
	e.ID = shortid.MustGenerate()
	e.Timestamp = time.Now()
}

// ActionEvent is an event that has a body.
type ActionEvent struct {
	Body interface{} `json:"body"`
}

// ActionEvent represents an action that took place.
type ClientActionEvent struct {
	Event
	ActionEvent
}

// GetBroadcastPayload will return the object to send to all chat users.
func (e *ClientActionEvent) GetPayload() EventPayload {
	return EventPayload{
		"id":        e.ID,
		"timestamp": e.Timestamp,
		"body":      e.Body,
		"type":      e.GetMessageType(),
	}
}

// GetMessageType will return the type of message.
func (e *ClientActionEvent) GetMessageType() EventType {
	return ClientAction
}

// SystemActionEvent is an event that represents an action that took place.
type SystemActionEvent struct {
	Event
	Message string `json:"message"`
	Status  int    `json:"status"`
}
