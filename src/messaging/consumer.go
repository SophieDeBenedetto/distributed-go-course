package messaging

import (
	"log"
	"github.com/streadway/amqp"
)

// Consumer receives messages from the queue
type Consumer struct {
	Server *Server
	Channel *amqp.Channel
	Queue  amqp.Queue
}

// NewConsumer returns a consumer struct
func NewConsumer(s *Server, queue string) *Consumer {
	ch, err := s.Conn.Channel()
	FailOnError(err, "Failed to open channel")
	q, err := ch.QueueDeclare(queue, false, false, false, false, nil)
	FailOnError(err, "Failed to declare and connect to queue")
	return &Consumer{
		Server: s,
		Queue:  q,
		Channel: ch,
	}
}

// Consume starts listening for messages from a queue
func (c *Consumer) Consume() {
	msgs, err := c.Channel.Consume(c.Queue.Name, "", true, false, false, false, nil)
	FailOnError(err, "Failed to start consumer")
	for msg := range msgs {
		log.Printf("Received message with body: %v", string(msg.Body))
	}
}

// Stop closes the channel connection
func (c *Consumer) Stop() {
	c.Channel.Close()
}
