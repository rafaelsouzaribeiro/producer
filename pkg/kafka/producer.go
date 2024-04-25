package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/segmentio/kafka-go"
)

func (brokers *Brokers) Send(ms *pkg.Message) {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers.Brokers,
		Topic:   ms.Topic,
	})

	writer := apmkafkago.WrapWriter(kWriter)

	m := kafka.Message{
		Value:   []byte(ms.Value),
		Headers: brokers.Headers,
	}

	err := writer.WriteMessages(context.Background(), m)

	if err != nil {
		fmt.Println("Error sending message:", err)
	}

}
