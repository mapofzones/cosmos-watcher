package rabbitmq

import (
	"encoding/json"
	"net/url"

	"github.com/attractor-spectrum/cosmos-watcher/x/tendermint-rabbit/tx"
	"github.com/streadway/amqp"
)

// TxQueue will call inside of itself another function which in an infinite loop recieves from channel
// same as websocket but the other way around
// send txs to first returned value and listen to errors from error channel
func TxQueue(nodeAddr url.URL) (chan<- []tx.Tx, <-chan error, error) {
	txs := make(chan []tx.Tx)
	errCh := make(chan error)
	conn, err := amqp.Dial(nodeAddr.String())
	if err != nil {
		return nil, nil, err
	}
	// channel handles API stuff for us
	ch, err := conn.Channel()
	// create query for our messages
	q, err := ch.QueueDeclare(
		"main",
		true,  // Durable means that messages are not lost when rabbitMQ exits
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return nil, nil, err
	}
	go func() {
		defer conn.Close()
		defer ch.Close()
		for {
			select {
			case data := <-txs:
				bytes, err := json.Marshal(data)
				if err != nil {
					errCh <- err
					close(txs)
					close(errCh)
					return
				}
				err = ch.Publish(
					"",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        bytes,
					},
				)
				if err != nil {
					errCh <- err
					close(txs)
					close(errCh)
					return
				}
			}
		}
	}()

	if err != nil {
		return nil, nil, err
	}
	return txs, errCh, nil
}
