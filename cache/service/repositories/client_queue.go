package repositories

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueClient struct {
	Connection *amqp.Connection
}

func NewQueueClient(user string, password string, host string, port int) *QueueClient {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d", user, password, host, port))
	if err != nil {
		log.Panic(err)
	}
	return &QueueClient{
		Connection: connection,
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (q *QueueClient) Consumer(exchange string, cacheclient CacheClient) {
	channel, err := q.Connection.Channel()
	failOnError(err, "failed to declare a new channel")

	queue, er := channel.QueueDeclare(
		"users_4",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(er, "failed to declare a new queue")

	er = channel.QueueBind(
		queue.Name,
		"#.#",
		exchange,
		false,
		nil,
	)
	failOnError(er, "failed to bind queue")

	msgs, er := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(er, "failed to consume queue")

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			keyslice := strings.Split(d.RoutingKey, ".")
			action := keyslice[1]
			userid, err := strconv.Atoi(keyslice[0])
			if err != nil {
				failOnError(err, "failed to convert user id to int")
			}
			fmt.Printf("Message received from %v, userid: %v, action: %v\n", queue.Name, userid, action)

			if action == "delete" {
				go func(userid int) {
					err := cacheclient.DeleteUserData(userid)
					if err != nil {
						failOnError(err, "failed to delete user from cache")
					}
				}(userid)
			}
		}
	}()
	<-forever
}
