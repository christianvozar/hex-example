package rabbitmq

import (
	"context"

	"github.com/christianvozar/hex-example/internal/domain"
	"github.com/streadway/amqp"
)

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewConsumer(amqpURI, queue string) (*Consumer, error) {
	conn, err := amqp.Dial(amqpURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Consumer{conn: conn, channel: ch, queue: queue}, nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan domain.Event, error) {
	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	events := make(chan domain.Event)

	go func() {
		defer close(events)

		for {
			select {
			case msg := <-msgs:
				events <- domain.Event{
					ID:      msg.MessageId,
					Payload: string(msg.Body),
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return events, nil
}
