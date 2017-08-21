// +build appengine

package rest

import (
	"net/http"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
)

func getContext(r *http.Request) context.Context {
	return context.WithValue(appengine.NewContext(r), ContextKeyEnvironment, "appengine")
}

// setNamespace sets a custom namespace if the `Namespace` variable is not nil
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

// Hostname returns the hostname of the current instance
func Hostname(ctx context.Context) (string, error) {
	return appengine.ModuleHostname(ctx, "", "", "")
}
