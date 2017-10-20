// +build !appengine,go1.7

package rest

import (
	"github.com/bradberger/context"
)

func setNamespace(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// NewTestContext returns a new context suitable for testing. Outside of appengine,
// this is just a fresh background context
func NewTestContext() (context.Context, func(), error) {
	return context.Background(), func() {}, nil
}
