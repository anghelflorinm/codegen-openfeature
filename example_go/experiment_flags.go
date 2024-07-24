package experimentflags

import (
	"codegen/providers"
	"context"

	"github.com/open-feature/go-sdk/openfeature"
)

// A flag for Open Feature.
var MyOpenFeatureFlag = struct {
	Value providers.BooleanProvider
}{
	Value: func(ctx context.Context, client *openfeature.Client) (bool, error) {
		return client.BooleanValue(ctx, "my_open_feature_flag", true, openfeature.EvaluationContext{})
	},
}
