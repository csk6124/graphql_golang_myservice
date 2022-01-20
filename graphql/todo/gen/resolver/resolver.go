package resolver

import "context"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ctx context.Context
}

// NewResolver new resolver reference
func NewResolver(
	ctx context.Context,
) Resolver {
	return Resolver{
		ctx: ctx,
	}
}
