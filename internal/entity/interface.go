package entity

type Iproducer interface {
	Send(queue Queue)
}
