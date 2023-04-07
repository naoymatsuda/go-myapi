package middlewares

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

func newTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo++
	mu.Unlock()

	return no
}

func SetTraceID(ctx context.Context, traceID int) context.Context{
	return context.WithValue(ctx, "traceID", traceID)
}

type traceIDKey struct{}

func GetTraceID(ctx context.Context) int {
type traceIDKey struct{}
id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}
