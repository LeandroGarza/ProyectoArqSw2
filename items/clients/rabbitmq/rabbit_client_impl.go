package rabbitmq

import (
	"strconv"
	"strings"

	dtos "items/dtos"
	"items/services/repositories"

	"context"
	"fmt"
	"log"
	"time"

	json "encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Channel *amqp.Channel
}

func NewRabbitmq(host string, port int) *RabbitMQ {
	portS := strconv.Itoa(port)
	dial := "amqp://user:password@" + host + ":" + portS + "/"
	conn, err := amqp.Dial(dial)
	if err != nil {
		panic(fmt.Sprintf("Error initializing RabbitMQ: %v", err))
	}

	ch, err := conn.Channel()
	if err != nil {
		panic(fmt.Sprintf("Error initializing RabbitMQ: %v", err))
	}

	fmt.Println("[RabbitMQ] Initialized connection")
	return &RabbitMQ{
		Channel: ch,
	}
}

func (queue RabbitMQ) PublishItems(ctx context.Context, items dtos.ItemsDto) error {
	q, err := queue.Channel.QueueDeclare(
		"newqueue", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body, err := json.Marshal(items)
	if err != nil {
		return err
	}

	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		}, // messahe
	)
	if err != nil {
		return err
	}

	log.Printf(" [RabbitMQ] Sent %s", body)

	return nil
}

func (queue RabbitMQ) PublishItem(ctx context.Context, item dtos.ItemDto) error {
	q, err := queue.Channel.QueueDeclare(
		"newqueue", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	body := item.Id
	err = queue.Channel.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // inmediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		}, //message
	)
	if err != nil {
		return err
	}

	log.Printf(" [RabbitMQ] Sent %s", body)

	return nil
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func (queue RabbitMQ) ConsumeUserUpdate(exchange string, ccache repositories.Repository, mongo repositories.Repository) {
	q, err := queue.Channel.QueueDeclare(
		"users_2", // name
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")

	err = queue.Channel.QueueBind(
		q.Name,
		"#.#",
		exchange,
		false,
		nil,
	)
	failOnError(err, "Failed to bind queue to exchange")

	msgs, err := queue.Channel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a consumer")

	forever := make(chan struct{})
	go func() {
		for d := range msgs {
			keyslice := strings.Split(d.RoutingKey, ".")
			action := keyslice[1]
			userid, err := strconv.Atoi(keyslice[0])
			if err != nil {
				failOnError(err, "failed to convert user id to int")
			}
			fmt.Printf("Message received from %v, userid: %v, action: %v\n", q.Name, userid, action)

			if action == "delete" {
				go func() {
					err := ccache.DeleteByUserId(context.TODO(), userid)
					if err != nil {
						failOnError(err, "failed to clear ccache")
					}
					er := mongo.DeleteByUserId(context.TODO(), userid)
					if er != nil {
						failOnError(err, "failed to delete item by userid en database")
					}
				}()
			}
		}
	}()
	log.Printf(" [*] Waiting for messages from %v. To exit press CTRL+C", q.Name)
	<-forever
}
