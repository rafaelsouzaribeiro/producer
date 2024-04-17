package producer

import (
	"github.com/rafaelsouzaribeiro/producer/internal/entity"
)

type ProducerUseCase struct {
	producer entity.Iproducer
}

func NewProducerUseCase(e entity.Iproducer) *ProducerUseCase {
	return &ProducerUseCase{
		producer: e,
	}
}
