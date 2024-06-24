package main

import (
	"log"

	"github.com/tolling/client"
)

//	type DistanceCalculator struct {
//		consumer DataConsumer
//	}

const (
	kafkaTopic         = "obudata"
	aggregatorEndpoint = "http://127.0.0.1:3000/aggregate"
)

// Tranport (http, grpc, kafka) -> attach business logic to this tranport

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc, client.NewClient(aggregatorEndpoint))
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
