package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"strconv"
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
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"task_queue",
		true,  // Durable: 队列是否持久化, false: 内存; true: 自带数据库
		false, // AutoDelete: 通常 AutoDelete = !Durable, 非自动删除队列
		false, // Exclusive: 当前链接独占的,私有的; 一般给false;
		false, // NoWait:true, 假定已经声明, 不允许其他链接修改; 通常给false;
		nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("Hello Work Queue: %s", strconv.Itoa(i))
		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "text/plain",
				Body:         []byte(data),
			},
		)
		if err != nil {
			return
		}
	}
}
