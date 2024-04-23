package kafka

import (
	"context"
	"fmt"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/segmentio/kafka-go"
)

func (brokers *Brokers) Send(ms *pkg.Message) {
	writerConfig := kafka.WriterConfig{
		Brokers: brokers.Brokers,
		Topic:   brokers.Topic,
	}

	writer := kafka.NewWriter(writerConfig)

	defer writer.Close()
	m := kafka.Message{
		Value: []byte(ms.Value),
	}

	err := writer.WriteMessages(context.Background(), m)

	if err != nil {
		fmt.Println("Error sending message:", err)
	}
}
