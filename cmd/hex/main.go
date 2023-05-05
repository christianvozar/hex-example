package main

import (
	"context"
	"log"

	"github.com/christianvozar/hex-example/internal/adapter/kafka"
	"github.com/christianvozar/hex-example/internal/adapter/postgres"
	"github.com/christianvozar/hex-example/internal/usecase"
)

func main() {
	ctx := context.Background()

	// Initialize the Postgres listener and Kafka consumer
	pgListener, err := postgres.NewListener()
	if err != nil {
		log.Fatalf("failed to create Postgres listener: %v", err)
	}

	kafkaConsumer, err := kafka.NewConsumer("your_topic_name")
	if err != nil {
		log.Fatalf("failed to create Kafka consumer: %v", err)
	}

	// Initialize the use case with the Postgres listener and Kafka consumer as dependencies
	watcher := usecase.NewWatcher(pgListener, kafkaConsumer)

	// Start watching for changes in the Postgres database
	if err := watcher.Watch(ctx); err != nil {
		log.Fatalf("failed to watch changes: %v", err)
	}
}
