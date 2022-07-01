package conn

import (
	"fmt"
	"os"
)

const (
	amqpConnStringTemplate = "amqp://%s:%s@%s:%s/"
)

func GetRabbitMQConnURL() string {
	host := os.Getenv("RABBITMQ_HOST")
	port := os.Getenv("RABBITMQ_PORT")
	user := os.Getenv("RABBITMQ_USER")
	password := os.Getenv("RABBITMQ_PASSWORD")
	return fmt.Sprintf(
		amqpConnStringTemplate,
		user,
		password,
		host,
		port,
	)
}
