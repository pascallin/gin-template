package events

type MessageEvent struct {
	ClientActionEvent
	Message int `json:"message"`
}

func (e *MessageEvent) GetPayload() EventPayload {
	return EventPayload{
		"type":      e.GetMessageType(),
		"id":        e.ID,
		"timestamp": e.Timestamp,
		"body":      e.Body,
	}
}

func (e *MessageEvent) GetMessageType() EventType {
	return ClientMessage
}
