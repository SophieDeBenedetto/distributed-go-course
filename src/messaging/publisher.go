package messaging

import "github.com/streadway/amqp"

// Publisher publishes to the queue
type Publisher struct {
	Server *Server
	Queue  amqp.Queue
}

// NewPublisher returns a publisher struct with the server and queue
func NewPublisher(s *Server, q amqp.Queue) *Publisher {
	return &Publisher{
		Server: s,
		Queue:  q,
	}
}

// Message returns a message to be published
func (p *Publisher) Message(contentType string, body []byte) amqp.Publishing {
	return amqp.Publishing{
		ContentType: contentType,
		Body:        body,
	}
}

// Publish publishes a message
func (p *Publisher) Publish(msg amqp.Publishing) {
	p.Server.Channel.Publish("", p.Queue.Name, false, false, msg)
}
