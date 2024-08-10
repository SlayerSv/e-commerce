package main

import (
	"context"

	"github.com/twmb/franz-go/pkg/kgo"
)

func (app *Application) NewKafkaClient() (*kgo.Client, error) {
	kafkaopts := []kgo.Opt{
		kgo.SeedBrokers("localhost:9092"),
		kgo.DefaultProduceTopic("errors"),
		kgo.ClientID("smartphones"),
	}
	kafka, err := kgo.NewClient(kafkaopts...)
	if err != nil {
		return nil, err
	}
	return kafka, nil
}

func (app *Application) NewErrorMessage(err error) {
	errmsg := &kgo.Record{
		Value: []byte(err.Error()),
		Topic: "errors",
	}
	ctx := context.Background()
	ferr := app.Kafka.ProduceSync(ctx, errmsg).FirstErr()
	if ferr != nil {
		app.ErrorLogger.Println(ferr)
	}
}
