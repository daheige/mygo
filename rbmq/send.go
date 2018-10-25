package main

import (
	"log"

	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://heige:zhuwei313@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.Publish(
		"Test_01", // exchange
		"Agent",   // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			MessageId:   time.Now().Format("2006-01-02 15:04:05"),
			Type:        "AgentJob",
			Body:        []byte("Hello world"),
		})
	log.Println("send ok")
	failOnError(err, "Failed to publish a message")
}
