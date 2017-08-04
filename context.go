package rest

// ContextKey is a string type for context value key names
type ContextKey string

// Context key location definitions
var (
	ContextKeyNamepace       ContextKey = "namespace"
	ContextKeyRequestBody    ContextKey = "request.body"
	ContextKeyRequestVars    ContextKey = "request.vars"
	ContextKeyResponseWriter ContextKey = "http.responsewriter"
	ContextKeyRequest        ContextKey = "request"
)
