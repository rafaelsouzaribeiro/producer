package di

import (
	"github.com/rafaelsouzaribeiro/producer/internal/infra/queue"
	"github.com/rafaelsouzaribeiro/producer/internal/usecase/producer"
)

func NewProducerUseCase() producer.ProducerUseCase {

	kafkaProducer := queue.NewKafkaProducer()
	myService := producer.NewProducerUseCase(kafkaProducer)

	return *myService

}
