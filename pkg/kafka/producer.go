package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func (brokers *Brokers) Send(ms *pkg.Message) {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers.Brokers,
		Topic:   ms.Topic,
	})

	writer := apmkafkago.WrapWriter(kWriter)

	m := kafka.Message{
		Value:   []byte(ms.Value),
		Headers: *brokers.GetHeader(ms),
	}

	err := writer.WriteMessages(context.Background(), m)

	if err != nil {
		fmt.Println("Error sending message:", err)
	}

}

func (c *Brokers) GetHeader(ms *pkg.Message) *[]protocol.Header {
	var headers []protocol.Header
	for _, header := range ms.Headers {
		headers = append(headers, protocol.Header{
			Key:   header.Key,
			Value: []byte(header.Value),
		})
	}

	return &headers
}
