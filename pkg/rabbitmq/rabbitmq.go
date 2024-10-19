package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	ch.Qos(100, 0, false)
	if err != nil {
		panic(err)
	}
	return ch, nil
}

//receive a channel (GO) and  write the msgs received from the queue to the channel
func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	//write the msgs received from the queue to the channel
	for msg := range msgs {
		out <- msg
	}
	return nil
}
