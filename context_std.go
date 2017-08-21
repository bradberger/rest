// +build !appengine

package rest

import (
	"net/http"
	"strings"

	"context"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(context.Background(), ContextKeyEnvironment, "standard")
}

// setNamespace sets a custom namespace if the `Namespace` variable is not nil
func setNamespace(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

// Hostname returns the hostname of the current instance
func Hostname(ctx context.Context) (string, error) {
	return strings.Split(Request(ctx).Host, ":")[0], nil
}
