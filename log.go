package rest

import "golang.org/x/net/context"

type LogFunc func(ctx context.Context, format string, args ...interface{})
