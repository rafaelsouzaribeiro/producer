package producer

func (c *ProducerUseCase) Producer(addrs []string, topic string, message []byte) {

	c.producer.Producer(addrs, topic, message)

}
