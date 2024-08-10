package main

import (
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func NewKafkaClient() (*kgo.Client, error) {
	kafkaopts := []kgo.Opt{
		kgo.SeedBrokers("localhost:9092"),
		kgo.DefaultProduceTopic("errors"),
		kgo.ConsumeTopics("errors"),
		kgo.ClientID("smartphones"),
	}
	kafka, err := kgo.NewClient(kafkaopts...)
	if err != nil {
		return nil, err
	}
	for {
		fetches := kafka.PollFetches(context.Background())
		if errs := fetches.Errors(); len(errs) > 0 {
			for err := range errs {
				log.Println(err)
			}
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			log.Println(record.Timestamp, string(record.Value), record.Topic)
		}
	}
}
