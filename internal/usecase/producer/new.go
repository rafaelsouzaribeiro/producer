package producer

import (
	"github.com/rafaelsouzaribeiro/producer/internal/ports"
)

type ProducerUseCase struct {
	producer ports.Iproducer
}

func NewProducerUseCase(e ports.Iproducer) *ProducerUseCase {
	return &ProducerUseCase{
		producer: e,
	}
}
