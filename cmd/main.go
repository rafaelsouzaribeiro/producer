package main

import (
	"sync"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	exec "github.com/rafaelsouzaribeiro/producer/pkg/kafka"
	"github.com/segmentio/kafka-go"
)

func main() {
	header := []kafka.Header{{Key: "your-header-key", Value: []byte("your-header-value")}}

	producer := exec.NewBrokers([]string{"springboot:9092"}, header)

	ms := pkg.Message{
		Value: "Testar",
		Topic: "contact-adm-insert",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		producer.Send(&ms)
		wg.Done()
	}()

	wg.Wait()

}
