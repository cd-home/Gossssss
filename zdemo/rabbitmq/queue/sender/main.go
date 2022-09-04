package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@10.211.55.18:5672")
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	defer ch.Close()
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
	}
	body := "Hello World 2!"
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		log.Println(err)
	}
}
