package main

import (
	"sync"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/rafaelsouzaribeiro/producer/pkg/kafka"
)

func main() {
	producer := kafka.NewBrokers([]string{"springboot:9092"})

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
