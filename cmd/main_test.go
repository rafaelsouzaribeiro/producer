package main

import (
	"strconv"
	"sync"
	"testing"

	"github.com/rafaelsouzaribeiro/producer/pkg"
	exec "github.com/rafaelsouzaribeiro/producer/pkg/kafka"
)

var count int = 0
var producer = exec.NewBrokers(&[]string{"springboot:9092"})
var wg sync.WaitGroup

func BenchmarkWriter(b *testing.B) {

	for i := 0; i < b.N; i++ {
		count++

		ms := pkg.Message{
			Value: "Test" + strconv.Itoa(count),
			Topic: "contact-adm-insert",
			Headers: []pkg.Header{
				{
					Key:   "your-header-key" + strconv.Itoa(count),
					Value: "your-header-value" + strconv.Itoa(count),
				},
				{
					Key:   "your-header-key" + strconv.Itoa(count),
					Value: "your-header-value" + strconv.Itoa(count),
				},
			},
		}

		wg.Add(1)

		go func() {
			producer.Send(&ms)
			wg.Done()
		}()
		wg.Wait()
	}

}
