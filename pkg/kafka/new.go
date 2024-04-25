package kafka

import "github.com/segmentio/kafka-go"

type Brokers struct {
	Brokers []string
	Headers []kafka.Header
}

func NewBrokers(brokers []string, header []kafka.Header) *Brokers {

	return &Brokers{
		Brokers: brokers,
		Headers: header,
	}
}
