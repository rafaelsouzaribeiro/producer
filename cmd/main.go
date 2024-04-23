package main

import (
	"github.com/rafaelsouzaribeiro/producer/pkg"
	"github.com/rafaelsouzaribeiro/producer/pkg/kafka"
)

func main() {
	producer := kafka.NewBrokers([]string{"springboot:9092"})

	ms := pkg.Message{
		Value: "Testar",
		Topic: "contact-adm-insert",
	}

	producer.Send(&ms)
}
