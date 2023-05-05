package usecase

import (
	"context"

	"github.com/christianvozar/hex-example/internal/domain"
)

type Watcher struct {
	listener domain.Listener
	consumer domain.Consumer
}

func NewWatcher(listener domain.Listener, consumer domain.Consumer) *Watcher {
	return &Watcher{
		listener: listener,
		consumer: consumer,
	}
}

func (w *Watcher) Watch(ctx context.Context) error {
	events, err := w.listener.Listen(ctx)
	if err != nil {
		return err
	}

	// Forward Postgres events to Kafka
	go func() {
		for event := range events {
			if err := w.consumer.Publish(event); err != nil {
				// Handle error or log it
			}
		}
	}()

	// Consume Kafka events and process them
	kafkaEvents, err := w.consumer.Consume(ctx)
	if err != nil {
		return err
	}

	for event := range kafkaEvents {
		// Process the event
	}

	return nil
}
