package main

import "log"

//	type DistanceCalculator struct {
//		consumer DataConsumer
//	}

const kafkaTopic = "obudata"

// Tranport (http, grpc, kafka) -> attach business logic to this tranport

func main() {
	var (
		err error
		svc CalculatorServicer
	)
	svc = NewCalculatorService()
	svc = NewLogMiddleware(svc)
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}