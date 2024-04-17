package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/rafaelsouzaribeiro/apache-kafka/producer"
	utils "github.com/rafaelsouzaribeiro/logs/cmd"
	"github.com/rafaelsouzaribeiro/producer/internal/entity"
)

type KafkaProducer struct {
}

func NewKafkaProducer() *KafkaProducer {
	return &KafkaProducer{}
}

func (c *KafkaProducer) Send(queues entity.Queue) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	produc := producer.NewProducer(queues.Addrs, queues.Topic,
		sarama.ByteEncoder(queues.Message), config, func(messages string) {
			utils.Insert(queues.Topic, messages, time.Now())
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
