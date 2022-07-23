# RabbitMq

### Introduction

RabbitMQ is a message broker: it accepts and forwards messages. You can think about it as a post office: when you put the mail that you want posting in a post box, you can be sure that the letter carrier will eventually deliver the mail to your recipient. In this analogy, RabbitMQ is a post box, a post office, and a letter carrier. 

RabbitMq is made up of the following components:

-  A Producer - This is simply a program that sends messages
-  A queue - this is the name for a post box which lives inside RabbitMQ. Although messages flow through RabbitMQ and your applications, they can only be stored inside a queue. A queue is only bound by the host's memory & disk limits, it's essentially a large message buffer. Many producers can send messages that go to one queue, and many consumers can try to receive data from one queue. This is how we represent a queue:

- Consumer -A consumer is just simply a program that mostly waits to receive messages

### Prerequisites
RabbitMq installed.U can install it from [here](https://www.rabbitmq.com/download.html)

### What we will do

We will write two simple go programs,one is a producer-a program that will send a message,and the other is a consumer program,one that will receive the message 
and print it.

