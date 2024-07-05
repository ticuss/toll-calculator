package client

import (
	"github.com/tolling/types"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	types.AggregatorClient
	Endpoint string
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := types.NewAggregatorClient(conn)
	return &GRPCClient{
		Endpoint:         endpoint,
		AggregatorClient: c,
	}, nil
}
