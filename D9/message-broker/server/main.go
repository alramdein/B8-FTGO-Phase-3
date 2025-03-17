package main

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Producer

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"batch8", // name
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	body := "Halo batch8! 123112312312312124124123"

	err = ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		panic(err)
	}
}
