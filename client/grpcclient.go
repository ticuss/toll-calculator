package client

import (
	"context"

	"github.com/tolling/types"
	"google.golang.org/grpc"
)

type GRPCClient struct {
	client   types.AggregatorClient
	Endpoint string
}

func NewGRPCClient(endpoint string) (*GRPCClient, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := types.NewAggregatorClient(conn)
	return &GRPCClient{
		Endpoint: endpoint,
		client:   c,
	}, nil
}

func (c *GRPCClient) Aggregate(ctx context.Context, aggReq *types.AggregateRequest) error {
	_, err := c.client.Aggregate(ctx, aggReq)
	return err
}
