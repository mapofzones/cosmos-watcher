package rabbitmq

import (
	"context"
	"log"

	codec "github.com/mapofzones/cosmos-watcher/pkg/codec"
	watcher "github.com/mapofzones/cosmos-watcher/pkg/types"
	"github.com/streadway/amqp"
	"github.com/tendermint/go-amino"
)

// BlockQueue will call inside of itself another function which in an infinite loop receives from channel
func BlockQueue(ctx context.Context, addr string, queue string) (chan<- watcher.Block, error) {
	log.Println("pkg.rabbitmq.rabbitmq.go - 1")
	blockStream := make(chan watcher.Block)
	conn, err := amqp.Dial(addr)
	if err != nil {
		log.Println("pkg.rabbitmq.rabbitmq.go - 2")
		return nil, err
	}

	log.Println("pkg.rabbitmq.rabbitmq.go - 3")
	// channel handles API stuff for us
	ch, err := conn.Channel()
	if err != nil {
		log.Println("pkg.rabbitmq.rabbitmq.go - 4")
		return nil, err
	}

	log.Println("pkg.rabbitmq.rabbitmq.go - 5")
	// confirm deliveries to ensure strict order
	if err := ch.Confirm(false); err != nil {
		log.Println("pkg.rabbitmq.rabbitmq.go - 6")
		return nil, err
	}

	log.Println("pkg.rabbitmq.rabbitmq.go - 7")
	notifications := make(chan amqp.Confirmation)
	ch.NotifyPublish(notifications)

	log.Println("pkg.rabbitmq.rabbitmq.go - 8")
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
		log.Println("pkg.rabbitmq.rabbitmq.go - 9")
		return nil, err
	}
	log.Println("pkg.rabbitmq.rabbitmq.go - 10")
	go func() {
		log.Println("pkg.rabbitmq.rabbitmq.go - 11")
		// initiate codec and register types for it to properly marshall data
		cdc := amino.NewCodec()
		codec.RegisterTypes(cdc)

		log.Println("pkg.rabbitmq.rabbitmq.go - 12")
		defer conn.Close()
		defer ch.Close()
		log.Println("pkg.rabbitmq.rabbitmq.go - 13")
		for {
			log.Println("pkg.rabbitmq.rabbitmq.go - 14")
			select {
			case block, ok := <-blockStream:
				log.Println("pkg.rabbitmq.rabbitmq.go - 15")
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
				log.Println("pkg.rabbitmq.rabbitmq.go - 16")
				// now wait for confirmation in order to preserve order
				select {
				case confirmation := <-notifications:
					log.Println("pkg.rabbitmq.rabbitmq.go - 17")
					if !confirmation.Ack {
						log.Println("pkg.rabbitmq.rabbitmq.go - 18")
						//server could not receive our  publishing
						close(blockStream)
						log.Println("pkg.rabbitmq.rabbitmq.go - 19")
						return
					}
				case <-ctx.Done():
					log.Println("pkg.rabbitmq.rabbitmq.go - 20")
					close(blockStream)
					log.Println("pkg.rabbitmq.rabbitmq.go - 21")
					return
				}

			case <-ctx.Done():
				log.Println("pkg.rabbitmq.rabbitmq.go - 22")
				close(blockStream)
				log.Println("pkg.rabbitmq.rabbitmq.go - 23")
				return
			}
			log.Println("pkg.rabbitmq.rabbitmq.go - 24")
		}
	}()

	log.Println("pkg.rabbitmq.rabbitmq.go - 25")
	return blockStream, nil
}
