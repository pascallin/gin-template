package events

// EventType is the type of a websocket event.
type EventType = string

const (
	ClientAction EventType = "CLIENT_ACTION"
	SystemAction EventType = "SYSTEM_ACTION"

	ClientMessage EventType = "CHAT_MESSAGE"
)
