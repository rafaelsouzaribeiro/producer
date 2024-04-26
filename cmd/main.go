package main

import (
	"sync"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	exec "github.com/rafaelsouzaribeiro/producer/pkg/kafka"
)

func main() {
	producer := exec.NewBrokers([]string{"springboot:9092"})

	ms := []pkg.Message{
		{
			Value: "Message 1",
			Topic: "contact-adm-isert",
			Headers: []pkg.Header{
				{
					Key:   "your-header-key1",
					Value: "your-header-value1",
				},
				{
					Key:   "your-header-key2",
					Value: "your-header-value2",
				},
			},
		},
		{
			Value: "Message 2",
			Topic: "contact-adm-isert",
			Headers: []pkg.Header{
				{
					Key:   "your-header-key1",
					Value: "your-header-value1",
				},
				{
					Key:   "your-header-key2",
					Value: "your-header-value2",
				},
			},
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		producer.Send(&ms)
		wg.Done()
	}()

	wg.Wait()

}
