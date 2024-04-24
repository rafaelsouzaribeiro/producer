package kafka

import (
	"context"
	"fmt"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/segmentio/kafka-go"
	apmkafkago "github.com/sohaibomr/apm-kafkago"
)

func (brokers *Brokers) Send(ms *pkg.Message) {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers.Brokers,
		Topic:   ms.Topic,
	})

	writer := apmkafkago.WrapWriter(kWriter)

	m := kafka.Message{
		Value: []byte(ms.Value),
	}

	err := writer.WriteMessages(context.Background(), m)

	if err != nil {
		fmt.Println("Error sending message:", err)
	}

}
