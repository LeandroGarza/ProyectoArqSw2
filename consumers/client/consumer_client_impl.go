package client

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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

func (q *QueueClient) ConsumeItems() {
	channel, err := q.Connection.Channel()
	if err != nil {
		failOnError(err, "failed to declare a new channel")
	}

	queue, err := channel.QueueDeclare(
		"newqueue",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "failed to declare a new queue")

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, fmt.Sprintf("failed to consume queue %v", queue.Name))

	for d := range msgs {
		go func(d amqp.Delivery) {
			resp, err := http.Post("http://localhost:8081/", "application/json", bytes.NewReader(d.Body))
			if err != nil {
				failOnError(err, fmt.Sprintf("failed to post message consumed on queue %v", queue.Name))
			}
			if resp.StatusCode != http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					failOnError(err, "failed to read response from post request")
				}
				fmt.Println(string(body))
			}
		}(d)
	}
}

func (q *QueueClient) ConsumeUserUpdates(exchange string, endpoints []string) {
	channel, err := q.Connection.Channel()
	if err != nil {
		failOnError(err, "failed to declare a new channel")
	}

	queue, err := channel.QueueDeclare(
		"user_consumer",
		false,
		false,
		false,
		false,
		nil,
	)

	failOnError(err, "failed to declare a new queue")

	err = channel.QueueBind(
		queue.Name,
		"#.#",
		exchange,
		false,
		nil,
	)
	failOnError(err, fmt.Sprintf("failed to bind queue %v to exchange %v", queue.Name, exchange))

	msgs, err := channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, fmt.Sprintf("failed to consume queue %v", queue.Name))

	var forever chan struct{}
	for d := range msgs {
		go func(delivery amqp.Delivery) {
			parts := strings.Split(delivery.RoutingKey, ".")
			action := parts[1]
			if action == "delete" {
				for _, endpoint := range endpoints {
					url := fmt.Sprintf("%s/%s", endpoint, string(delivery.Body))

					req, err := http.NewRequest(http.MethodDelete, url, nil)
					failOnError(err, fmt.Sprintf("failed to create DELETE request for endpoint: %s", url))

					// realizar la solicitud DELETE
					client := http.Client{}
					resp, err := client.Do(req)
					failOnError(err, fmt.Sprintf("failed to execute DELETE request for endpoint: %s", url))

					defer resp.Body.Close()
				}
			}
			fmt.Printf("Message, recieved from %v: %v", queue.Name, string(delivery.Body))
		}(d)
	}
	log.Printf(" [*] Waiting for messages from %v.", queue.Name)
	<-forever
}
