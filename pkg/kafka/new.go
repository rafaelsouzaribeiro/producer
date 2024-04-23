package kafka

type Brokers struct {
	Brokers []string
	Topic   string
}

func NewBrokers(brokers []string, topic string) *Brokers {

	return &Brokers{
		Brokers: brokers,
		Topic:   topic,
	}
}
