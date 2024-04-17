package producer

import "github.com/rafaelsouzaribeiro/producer/internal/entity"

func (c *ProducerUseCase) Send(addrs []string, topic string, message []byte) {

	queue := entity.NewQueue(addrs, topic, message)
	c.producer.Send(*queue)

}
