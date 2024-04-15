package kafka

import (
	"github.com/rafaelsouzaribeiro/producer/internal/infra/queue/kafka"
	"github.com/rafaelsouzaribeiro/producer/internal/usecase/producer"
)

func NewProducerUseCase() producer.ProducerUseCase {

	kafkaProducer := kafka.NewKafkaProducer()
	myService := producer.NewProducerUseCase(kafkaProducer)

	return *myService

}
