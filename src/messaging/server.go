package messaging

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Server the RabbitMQ server
type Server struct {
	RabbitMQUsername string
	RabbitMQPassword string
	RabbitMQHost     string
	Conn             *amqp.Connection
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
	FailOnError(err, "Failed to connect to RabbitMQ")
	s.Conn = conn
}

// Close closes the connection and channel
func (s *Server) Close() {
	s.Conn.Close()
}
