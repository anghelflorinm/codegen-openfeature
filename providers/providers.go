package providers

import (
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

type BooleanProvider func(ctx context.Context, client *openfeature.Client) (bool, error)
type FloatProvider func(ctx context.Context, client *openfeature.Client) (float64, error)
type IntProvider func(ctx context.Context, client *openfeature.Client) (int, error)
type StringProvider func(ctx context.Context, client *openfeature.Client) (string, error)
