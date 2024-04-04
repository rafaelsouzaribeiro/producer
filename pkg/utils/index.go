package utils

import "github.com/rafaelsouzaribeiro/producer/internal/infra/di"

func Producer(addrs []string, topic string, message []byte) {
	producer := di.NewProducerUseCase()

	producer.Producer(addrs, topic, message)

}
