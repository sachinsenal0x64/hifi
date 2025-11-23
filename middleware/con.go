package middleware

import (
	"context"
	"fmt"
)

func Con() (context.Context, *Router, error) {

	ctx := context.Background()
	store, err := NewRouter("127.0.0.1:6379")

	if err != nil {
		return ctx, nil, fmt.Errorf("failed to create router: %w", err)
	}

	if store == nil {
		return ctx, nil, fmt.Errorf("router is nil after initialization")
	}

	return ctx, store, nil
}
