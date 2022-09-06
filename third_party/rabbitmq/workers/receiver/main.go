package main

import (
	"github.com/streadway/amqp"
	"log"
	"time"
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
	q, err := ch.QueueDeclare(
		"task_queue",
		true,  // Durable: 队列是否持久化, false: 内存; true: 自带数据库
		false, // AutoDelete: 通常 AutoDelete = !Durable, 非自动删除队列
		false, // Exclusive: 当前链接独占的,私有的; 一般给false;
		false, // NoWait:true, 假定已经声明, 不允许其他链接修改; 通常给false;
		nil,
	)
	err = ch.Qos(
		1, // prefetch count, 消息最多一个, 必须等待前一个确认, worker 的公平调度
		0,
		false,
	)
	if err != nil {
		log.Println(err)
	}
	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false, // autoAck = false, 采用手动确认方式
		false,
		false, nil,
	)
	if err != nil {
		log.Println(err)
		return
	}
	for msg := range msgs {
		log.Printf("Receiver: %s", msg.Body)
		time.Sleep(time.Second)
		log.Println("Done.")
		// 处理完任务采用手动确认
		err := msg.Ack(false)
		if err != nil {
			log.Println(err)
		}
	}
}
