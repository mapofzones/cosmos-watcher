package rabbitmq

import (
	"context"

	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/streadway/amqp"
	"github.com/tendermint/go-amino"
)

// BlockQueue will call inside of itself another function which in an infinite loop receives from channel
func BlockQueue(ctx context.Context, addr string, queue string, cdc *amino.Codec) (chan<- watcher.Block, error) {
	blockStream := make(chan watcher.Block)
	conn, err := amqp.Dial(addr)
	if err != nil {
		return nil, err
	}

	// channel handles API stuff for us
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	// confirm deliveries to ensure strict order
	if err := ch.Confirm(false); err != nil {
		return nil, err
	}

	notifications := make(chan amqp.Confirmation)
	ch.NotifyPublish(notifications)

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
		return nil, err
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
				data := cdc.MustMarshalJSON(block)
				err := ch.Publish(
					"",
					q.Name,
					true,
					false,
					amqp.Publishing{
						ContentType: "application/json",
						Body:        data,
					},
				)
				if err != nil {
					close(blockStream)
					return
				}
				// now wait for confirmation in order to preserve order
				select {
				case confirmation := <-notifications:
					if !confirmation.Ack {
						//server could not receive our  publishing
						close(blockStream)
						return
					}
				case <-ctx.Done():
					close(blockStream)
					return
				}

			case <-ctx.Done():
				close(blockStream)
				return
			}
		}
	}()

	return blockStream, nil
}
