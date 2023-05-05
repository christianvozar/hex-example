package kafka

import (
	"context"

	"github.com/christianvozar/hex-example/internal/domain"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
	topic    string
}

func NewConsumer(topic string) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "your_kafka_bootstrap_servers",
		"group.id":          "your_group_id",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{consumer: c, topic: topic}, nil
}

func (c *Consumer) Consume(ctx context.Context) (<-chan domain.Event, error) {
	if err := c.consumer.Subscribe(c.topic, nil); err != nil {
		return nil, err
	}

	events := make(chan domain.Event)

	go func() {
		defer close(events)

		for {
			msg, err := c.consumer.ReadMessage(ctx.Done())
			if err != nil {
				return
			}

			events <- domain.Event{
				ID:      string(msg.TopicPartition.Topic),
				Payload: string(msg.Value),
			}
		}
	}()

	return events, nil
}
