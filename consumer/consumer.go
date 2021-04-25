package consumer

import "github.com/segmentio/kafka-go"

type Consumer interface {
	Consume() error
}

type kafkaConsumer struct {
}

func NewKafkaConsumer() Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MinBytes: 1,    // 1B
		MaxBytes: 10e6, // 10MB
	})
	return &kafkaConsumer{}
}

func (c *kafkaConsumer) Consume() error {
	return nil
}
