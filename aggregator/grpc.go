package main

import (
	"context"

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

func (s *GRPCSAggregatorerver) Aggregate(ctx context.Context, req *types.AggregateRequest) (*types.None, error) {
	distance := types.Distance{
		OBUID: int(req.ObuID),
		Value: req.Value,
		Unix:  req.Unix,
	}
	return &types.None{}, s.svc.AggregateDistance(distance)
}
