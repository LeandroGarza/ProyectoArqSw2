package clients

/*
package clients

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func ConsumeItems() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"newqueue", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			resp, err := http.Post("http://localhost:8983/solr/items/update/json/docs?commit=true", "application/json", bytes.NewReader(d.Body))
			if err != nil {
				log.Println(err)
			} else {
				go func() {
					defer resp.Body.Close()

					responseBody, err := ioutil.ReadAll(resp.Body)
					if err != nil {
						log.Println(err)
					} else {
						fmt.Println(string(responseBody))
					}
				}()
			}

			fmt.Printf("Message, recieved from %v: %v", q.Name, string(d.Body))
		}
	}()

	log.Printf(" [*] Waiting for messages from %v.", q.Name)
	<-forever
}

func ConsumeUserUpdates() {
	conn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		"users_1", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		"users_1", // name
		"#.#",     // routing key
		"users",   // exchange name
		false,     //
		nil,       //
	)
	failOnError(err, "Failed to bind queue to exchange")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			fmt.Printf("Message received from %v, Id.action: %v", q.Name, string(d.RoutingKey))
		}
	}()

	log.Printf(" [*] Waiting for messages from %v. To exit press CTRL+C", q.Name)
	<-forever
}
*/
