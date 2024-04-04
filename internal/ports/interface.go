package ports

type Iproducer interface {
	Producer(addrs []string, topic string, message []byte)
}
