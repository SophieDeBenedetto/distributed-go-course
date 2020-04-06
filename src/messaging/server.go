package messaging

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// Server the RabbitMQ server
type Server struct {
	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	Conn             *amqp.Connection
	Channel          *amqp.Channel
}

// NewRabbitMQServer returns the new rabbitmq server with the connection and channel
func NewRabbitMQServer(username string, password string, host string) *Server {
	return &Server{
		RabbitMQUsername: username,
		RabbitMQPassword: password,
		RabbitMQHost:     host,
	}
}

// Connect connects to the RabbitMQ server
func (s *Server) Connect() {
	connectionAddr := fmt.Sprintf("amqp://%s:%s@%s", s.RabbitMQUsername, s.RabbitMQPassword, s.RabbitMQHost)
	conn, err := amqp.Dial(connectionAddr)
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open channel")
	s.Conn = conn
	s.Channel = ch
}

// DeclareQueue declares the queue
func (s *Server) DeclareQueue(queue string) amqp.Queue {
	q, err := s.Channel.QueueDeclare(queue, false, false, false, false, nil)
	failOnError(err, "Failed to declare and connect to queue")
	return q
}

// Close closes the connection and channel
func (s *Server) Close() {
	s.Channel.Close()
	s.Conn.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(err)
	}
}
