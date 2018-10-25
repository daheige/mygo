package main

import (
	"log"
	"sync"

	"encoding/json"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type msgInfo struct {
	Id   int
	Name string
}

// 只能在安装 rabbitmq 的服务器上操作
func main() {
	conn, err := amqp.Dial("amqp://heige:zhuwei313@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			msg := &msgInfo{
				Id:   i,
				Name: body,
			}

			json_data, _ := json.Marshal(msg)

			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        json_data,
				})

			if err != nil {
				log.Println(err, "Failed to publish a message")
				continue
			}

			log.Printf(" [x] Sent %s", string(json_data))
		}
	}()
	wg.Wait()

	log.Println("All msg has send")
}
