package main

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tolling/types"
)

type LogMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (m *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		fmt.Println(distance)
		logrus.WithFields(
			logrus.Fields{
				"took": time.Since(start),
				"err":  err,
			}).Info("AggregateDistance")
	}(time.Now())
	err = m.next.AggregateDistance(distance)
	return
}

func (m *LogMiddleware) CalculateInvoice(obyID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)
		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":     time.Since(start),
			"err":      err,
			"obuID":    obyID,
			"distance": distance,
			"amount":   amount,
		}).Info("CalculateInvoice")
	}(time.Now())
	inv, err = m.next.CalculateInvoice(obyID)
	return
}
