package server

import (
	"encoding/json"
	"sync"

	"github.com/sirupsen/logrus"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	mu sync.RWMutex
	// Registered clients.
	clients map[*Client]bool

	// send outbound message payload to all clients
	broadcast chan []byte

	// receive inbound message payload from all clients
	inbound chan clientEvent

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		inbound:    make(chan clientEvent),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) eventReceived(event clientEvent) {
	var typecheck map[string]interface{}
	if err := json.Unmarshal(event.data, &typecheck); err != nil {
		logrus.Debugln(err)
	}

	eventType := typecheck["type"]

	logrus.WithFields(logrus.Fields{
		"typecheck": typecheck,
		"eventType": eventType,
	}).Debug("eventReceived")
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			logrus.Info("Client register")
			h.clients[client] = true
			go client.writePump()
		case client := <-h.unregister:
			logrus.Info("Client unregister")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					h.mu.Lock()
					close(client.send)
					delete(h.clients, client)
					h.mu.Unlock()
				}
			}
		case message := <-h.inbound:
			h.eventReceived(message)
		}
	}
}
