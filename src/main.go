package main

import (
	"fmt"

	"github.com/m/src/messaging"
)

func main() {
	fmt.Println("Running...")
	rabbitServer := messaging.NewRabbitMQServer("guest", "guest", "localhost:5672")
	defer rabbitServer.Close()
	rabbitServer.Connect()
	q := rabbitServer.DeclareQueue("hello")
	publisher := messaging.NewPublisher(rabbitServer, q)
	publisher.Publish(publisher.Message("text/plain", []byte("Hello World")))
}
