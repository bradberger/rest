package rest

import (
	"net/http"

	"golang.org/x/net/context"
)

// ContextKey is a string type for context value key names
type ContextKey string

// Context key location definitions
var (
	ContextKeyNamepace       ContextKey = "namespace"
	ContextKeyRequestBody    ContextKey = "request.body"
	ContextKeyRequestVars    ContextKey = "request.vars"
	ContextKeyResponseWriter ContextKey = "http.responsewriter"
	ContextKeyRequest        ContextKey = "request"
	ContextKeyResponseCode   ContextKey = "response.code"
	ContextKeyInitialized    ContextKey = "initialized"
	ContextKeyResponseBody   ContextKey = "response.body"
	ContextKeyEnvironment    ContextKey = "environment"
)

func SetCode(ctx context.Context, code int) context.Context {
	return context.WithValue(ctx, ContextKeyResponseCode, code)
}

func GetCode(ctx context.Context) int {
	if v := ctx.Value(ContextKeyResponseCode); v != nil {
		return v.(int)
	}
	return http.StatusOK
}
