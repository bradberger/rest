// +build !appengine

package rest

import (
	"github.com/bradberger/context"
)

func setNamespace(ctx context.Context) (context.Context, error) {
	return ctx, nil
}
