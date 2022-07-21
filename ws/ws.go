package ws

import (
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

var _hub *Hub

func Start() error {
	_hub = NewHub()
	go _hub.run()
	return nil
}

// serveWs handles websocket requests from the peer.
func ServeWs(c *gin.Context) {
	// upgrade http to ws
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Error(err)
		return
	}

	// register ws client
	client := &Client{
		id:   generateSegmentId(),
		hub:  _hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.handler()
}

func generateSegmentId() string {
	sid, err := shortid.New(1, strings.Replace(shortid.DefaultABC, "-", "=", -1), 2342)
	if err != nil {
		log.Error(err)
		return ""
	}
	return sid.MustGenerate()
}

// GetClientsForUser will return chat connections that are owned by a specific user.
func FindClientByID(id string) (*Client, bool) {
	client, found := _hub.clients[id]
	return client, found
}

func GetWebsocketConnectionCount() int {
	return len(_hub.clients)
}
