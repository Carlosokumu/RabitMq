# RabbitMq

### Introduction

RabbitMQ is a message broker: it accepts and forwards messages. You can think about it as a post office: when you put the mail that you want posting in a post box, you can be sure that the letter carrier will eventually deliver the mail to your recipient. In this analogy, RabbitMQ is a post box, a post office, and a letter carrier. 

RabbitMq is made up of the following components:

-  A Producer - This is simply a program that sends messages
-  A queue - this is the name for a post box which lives inside RabbitMQ. Although messages flow through RabbitMQ and your applications, they can only be stored inside a queue. A queue is only bound by the host's memory & disk limits, it's essentially a large message buffer. Many producers can send messages that go to one queue, and many consumers can try to receive data from one queue. This is how we represent a queue:

- Consumer -A consumer is just simply a program that mostly waits to receive messages

### Prerequisites
U will need RabbitMq installed on your machine.U can install it from [here](https://www.rabbitmq.com/download.html)

### What we will do

We will write two simple go programs,one is a producer-a program that will send a message,and the other is a consumer program,one that will receive the message 
and print it.

In the diagram below, "P" is our producer and "C" is our consumer. The box in the middle is a queue - a message buffer that RabbitMQ keeps on behalf of the consumer.

![RabbitMq](https://github.com/Carlosokumu/RabitMq/blob/master/images/python-one.png)

#### The Go RabbitMQ client library
There are a number of clients for RabbitMQ in many different languages. We'll use the Go amqp client in this app.
So go ahead and install it:

`go get github.com/rabbitmq/amqp091-go`

##### Sending
 Our sender will be in `send.go` file.The sender will connect to RabbitMQ,send a single message, then exit.
 We first import the library
 
 ```
 package main

import (
   "log"
    amqp "github.com/rabbitmq/amqp091-go"
)
```
We also need a helper function to check the return value for each amqp connection call:

```
func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}
```
And we connect to the RabbitMq server and check for possible errors with the helper function we created above

```
conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
failOnError(err, "Failed to connect to RabbitMQ")
defer conn.Close()

```
The connection abstracts the socket connection, and takes care of protocol version negotiation and authentication and so on for us. Next we create a channel using the the connection, which is where most of the API for getting things done resides:

```
 ch, err := conn.Channel()
 failOnError(err, "Failed to open a channel")
 defer ch.Close()

```
To send, we must declare a queue for us to send to; then we can publish a message to the queue:

```
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
err = ch.Publish(
  "",     // exchange
  q.Name, // routing key
  false,  // mandatory
  false,  // immediate
  amqp.Publishing {
    ContentType: "text/plain",
    Body:        []byte(body),
  })
failOnError(err, "Failed to publish a message")
log.Printf(" [x] Sent %s\n", body)

```
##### Receiving
