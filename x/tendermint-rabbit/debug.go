package watcher

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/streadway/amqp"
)

type DebugData struct {
	Data  []byte
	Chain string
	Time  time.Time
}

func (d DebugData) Marshal() []byte {
	bytes, _ := json.Marshal(d)
	return bytes
}

func DebugSend(rabbit url.URL, data DebugData) {
	conn, _ := amqp.Dial(rabbit.String())
	defer conn.Close()
	// channel handles API stuff for us
	ch, _ := conn.Channel()
	defer conn.Close()
	q, _ := ch.QueueDeclare(
		"debug",
		true,  // Durable means that messages are not lost when rabbitMQ exits
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data.Marshal(),
		},
	)
}
