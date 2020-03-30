package rabbitmq

// // create func which in an infinite loop recieves from channel
// // same as websocket but the other way around
// func (l Listener) sample() (chan<- Tx, error) {
// 	conn, err := amqp.Dial(l.rabbitMQAddr.String())
// 	if err != nil {
// 		l.logger.Println("failed to connect to RabbitMQ")
// 		return nil, err
// 	}
// 	// channel handles API stuff for us
// 	ch, err := conn.Channel()
// 	// create query for our messages
// 	q, err := ch.QueueDeclare(
// 		"test",
// 		true,  // Durable means that messages are not lost when rabbitMQ exits
// 		false, // auto-delete
// 		false, // exclusive
// 		false, // no-wait
// 		nil,
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// send the message
// 	err = ch.Publish(
// 		"",     // default exchange
// 		q.Name, // key to find queue to which send messages
// 		false,  // mandatory
// 		false,  // immediate
// 		amqp.Publishing{
// 			ContentType: "text/plain",
// 			Body:        []byte("txs"),
// 		})

// 	if err != nil {
// 		return nil, err
// 	}
// 	defer conn.Close()
// 	defer ch.Close()
// 	return nil, nil
// }
