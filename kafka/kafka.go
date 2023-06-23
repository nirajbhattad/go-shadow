package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consumer() {
	brokers := []string{"localhost:9092"}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       "sample_topic",
		GroupID:     "my_consumer_group",
		StartOffset: kafka.FirstOffset,
	})

	// Consume messages
	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Failed to read message:", err)
		}

		fmt.Println("Received message:", string(message.Value))
	}
}
