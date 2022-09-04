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
	if err != nil {
		log.Println(err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		log.Println(err)
	}
	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Wait.....")
	for msg := range msgs {
		log.Printf("Recevied Message: %s", msg.Body)
	}
}
