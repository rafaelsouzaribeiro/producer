package kafka

import (
	"context"
	"fmt"

	apmkafkago "github.com/rafaelsouzaribeiro/apm-kafkago/pkg"
	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

func (brokers *Brokers) Send(messages *[]pkg.Message) {
	if len(*messages) == 0 {
		fmt.Println("No messages to send")
		return
	}

	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers.Brokers,
		Topic:   (*messages)[0].Topic,
	})

	writer := apmkafkago.WrapWriter(kWriter)

	var kafkaMessages []kafka.Message
	for _, ms := range *messages {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Value:   []byte(ms.Value),
			Headers: *brokers.GetHeader(&ms),
		})
	}

	err := writer.WriteMessages(context.Background(), kafkaMessages...)

	if err != nil {
		fmt.Println("Error sending messages:", err)
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
