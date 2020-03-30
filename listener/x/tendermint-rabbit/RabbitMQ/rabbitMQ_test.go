package rabbitmq

import (
	"fmt"
	"testing"

	"github.com/streadway/amqp"
)

const defaultAddr = "amqp://guest:guest@localhost:5672/"

func TestConnectionSend(t *testing.T) {
	conn, err := amqp.Dial(defaultAddr)
	if err != nil {
		fmt.Printf("failed to connect to RabbitMQ, %s", err)
		t.FailNow()
	}
	// channel handles API stuff for us
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	// create query for our messages
	q, err := ch.QueueDeclare(
		"test",
		true,  // Durable means that messages are not lost when rabbitMQ exits
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	// send the message
	err = ch.Publish(
		"",     // default exchange
		q.Name, // key to find queue to which send messages
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("txs"),
		})

	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer conn.Close()
	defer ch.Close()
}

func TestConnectionRecieve(t *testing.T) {
	conn, err := amqp.Dial(defaultAddr)
	if err != nil {
		fmt.Printf("failed to connect to RabbitMQ, %s", err)
		t.FailNow()
	}
	// channel handles API stuff for us
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	// create query for our messages
	q, err := ch.QueueDeclare(
		"test",
		true,  // Durable means that messages are not lost when rabbitMQ exits
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer name
		false, // auto acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	for m := range msgs {
		fmt.Println(string(m.Body))
	}
}
