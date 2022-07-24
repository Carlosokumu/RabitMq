package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	_, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}
}
