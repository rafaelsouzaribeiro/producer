package kafka

type Brokers struct {
	Brokers []string
}

func NewBrokers(brokers []string) *Brokers {

	return &Brokers{
		Brokers: brokers,
	}
}
