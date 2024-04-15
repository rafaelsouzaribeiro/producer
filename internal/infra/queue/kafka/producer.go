package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/producer"
	utils "github.com/rafaelsouzaribeiro/logs/cmd"
)

type KafkaProducer struct {
	// Aqui você pode adicionar quaisquer outras dependências que sua implementação precise
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{}
}

func (c *KafkaProducer) Producer(addrs []string, topic string, message []byte) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	produc := producer.NewProducer(addrs, topic,
		sarama.ByteEncoder(message), config, func(messages string) {
			utils.Insert(topic, messages, time.Now())
		})
	prod, err := produc.GetProducer()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := (*prod).Close(); err != nil {
			panic(err)
		}
	}()

	produc.SendMessage(prod)

}
