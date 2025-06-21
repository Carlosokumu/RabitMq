package main

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	var wg sync.WaitGroup
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	log.Println("Connected successfully!")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel from connection: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare queque from channel")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Consume for ch: %s failed with error: %v", q.Name, err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	wg.Wait()
}
