package sender

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pascallin/gin-template/conn"
	"github.com/streadway/amqp"
)

var (
	QUEUE_NAME = "hello"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func init() {
	// load .env
	godotenv.Load()
}

// @Summary send mq message
// @Description send mq message
// @Tags mq
// @Accept  json
// @Produce json
// @Router /mq [post]
func SendHelloRoute(c *gin.Context) {
	SendHello("hello pascal!")
}

func SendHello(body string) {

	conn, err := amqp.Dial(conn.GetRabbitMQConnURL())
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		QUEUE_NAME, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}
