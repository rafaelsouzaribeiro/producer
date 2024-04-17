package entity

type Queue struct {
	Addrs   []string
	Topic   string
	Message []byte
}

func NewQueue(Addrs []string, Topic string, Message []byte) *Queue {

	return &Queue{
		Addrs:   Addrs,
		Topic:   Topic,
		Message: Message,
	}
}
