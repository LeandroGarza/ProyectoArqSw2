package repositories

import (
	"fmt"
	"log"
	"time"

	"context"
	e "users/utils/errors"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueClientImpl struct {
	Connection *amqp.Connection
}

func NewQueueClientImpl(user string, password string, host string, port int) *QueueClientImpl {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port))
	if err != nil {
		log.Panic(err)
	}

	return &QueueClientImpl{
		Connection: connection,
	}
}

func (q *QueueClientImpl) SendMessage(userid int, action string, message string) e.ApiError {
	channel, err := q.Connection.Channel()
	if err != nil {
		return e.NewBadRequestApiError("Failed to open a channel")
	}

	err = channel.ExchangeDeclare(
		"users", // name
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return e.NewBadRequestApiError("Failed to declare an exchange")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := message
	err = channel.PublishWithContext(
		ctx,                                  // context
		"users",                              // exchange
		fmt.Sprintf("%d.%s", userid, action), // key
		false,                                // mandatory
		false,                                // inmediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}) // message
	if err != nil {
		return e.NewBadRequestApiError("Failed to publish a message")
	}
	log.Printf("[X] Sent %s\n", body)
	return nil
}
