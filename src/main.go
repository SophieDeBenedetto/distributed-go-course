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

	publisher := messaging.NewPublisher(rabbitServer, "hello")
	defer publisher.Stop()
	go func() {
		for {
			publisher.Publish(publisher.Message("text/plain", []byte("Hello World")))
		}
	}()

	consumer := messaging.NewConsumer(rabbitServer, "hello")
	defer consumer.Stop()
	go func() {
		consumer.Consume()
	}()
	var a string
	fmt.Scanln(&a)
}
