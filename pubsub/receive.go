package pubsub

import (
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-template/conn"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

var (
	QUEUE_NAME = "hello"
)

func init() {
	godotenv.Load()
}

func Listen() {

	conn, err := amqp.Dial(conn.GetRabbitMQConnURL())
	if err != nil {
		log.Error("%s: %s", err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Error("%s: %s", err, "Failed to open a channel")
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Error("%s: %s", err, "Failed to declare a queue")
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Error("%s: %s", err, "Failed to register a consumer")
		return
	}

	for d := range msgs {
		log.Infof("Received a message: %s", d.Body)
	}

}
