// +build appengine

package rest

import (
	"github.com/bradberger/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/aetest"
)

func setNamespace(ctx context.Context) (context.Context, error) {
	if Namespace == nil {
		return ctx, nil
	}
	ns, err := Namespace(ctx)
	if err != nil {
		return ctx, err
	}
	return appengine.Namespace(ctx, ns)
}

// NewTestContext returns a new context suitable for testing. In the appengine
// environment, this is an aetest context
func NewTestContext() (context.Context, func(), error) {
	return aetest.NewContext()
}
