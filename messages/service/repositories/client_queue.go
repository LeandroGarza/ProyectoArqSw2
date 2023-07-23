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

func (q *QueueClient) ConsumeUserUpdates(exchange string, messageclient MessageClient) {
	ch, err := q.Connection.Channel()
	failOnError(err, "Failed to open a channel")

	queue, err := ch.QueueDeclare(
		"users_3", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		queue.Name, // name
		"#.#",      // routing key
		exchange,   // exchange name
		false,      //
		nil,        //
	)
	failOnError(err, "Failed to bind queue to exchange")

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

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
					er := messageclient.DeleteMessagesByUser(userid)
					if er != nil {
						failOnError(er, "Failed to delete user")
					}
				}(userid)
			}
		}
	}()

	log.Printf(" [*] Waiting for messages from %v. To exit press CTRL+C\n", queue.Name)
	<-forever
}
