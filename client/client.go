package client

import (
	"context"

	"github.com/tolling/types"
)

type Client interface {
	Aggregate(context.Context, *types.AggregateRequest) error
}
