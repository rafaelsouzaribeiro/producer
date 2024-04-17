package cmd

import "github.com/rafaelsouzaribeiro/producer/internal/infra/di/kafka"

func Producer(addrs []string, topic string, message []byte) {
	producer := kafka.NewProducerUseCase()
	producer.Send(addrs, topic, message)

}
