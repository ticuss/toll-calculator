package main

import (
	"github.com/tolling/types"
)

type GRPCSAggregatorerver struct {
	types.UnimplementedAggregatorServer
	svc Aggregator
}

func NewAggregatorGRPCServer(svc Aggregator) *GRPCSAggregatorerver {
	return &GRPCSAggregatorerver{
		svc: svc,
	}
}

func (s *GRPCSAggregatorerver) AggregateDistance(req *types.AggregateRequest) error {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return s.svc.AggregateDistance(distance)
}
