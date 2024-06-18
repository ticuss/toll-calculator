package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tolling/types"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{next: next}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {
	defer func(start time.Time) {
		// Do something with the data
		logrus.WithFields(logrus.Fields{
			"obu_id":    data.OBUID,
			"latitude":  data.Lat,
			"longitude": data.Long,
			"took":      time.Since(start),
		}).Info("Received data")
	}(time.Now())
	return l.next.ProduceData(data)
}
