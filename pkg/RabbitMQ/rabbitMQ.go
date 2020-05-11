package rabbitmq

import (
	"context"

	block "github.com/mapofzones/cosmos-watcher/pkg/block/types"
	"github.com/streadway/amqp"
)

// BlockQueue will call inside of itself another function which in an infinite loop receives from channel
// same as websocket but the other way around
// send blocks to blocks channel value and listen to errors from error channel
func BlockQueue(ctx context.Context, addr string, queue string) (chan<- block.Block, <-chan error, error) {
	blockStream := make(chan block.Block)
	errCh := make(chan error)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, nil, err
	}

	// channel handles API stuff for us
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	// create query for our messages
	q, err := ch.QueueDeclare(
		queue,
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
			case block, ok := <-blockStream:
				if !ok {
					return
				}
				err = ch.Publish(
					"",
					q.Name,
					false,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        block.JSON(),
					},
				)
				if err != nil {
					errCh <- err
					close(blockStream)
					close(errCh)
					return
				}
			case <-ctx.Done():
				close(blockStream)
				close(errCh)
				return
			}
		}
	}()

	return blockStream, errCh, nil
}
