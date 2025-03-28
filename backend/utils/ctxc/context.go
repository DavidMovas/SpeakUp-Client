package ctxc

import (
	"context"
	"time"
)

const shortContextTimeout = time.Second * 5

func ShortContext() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), shortContextTimeout)

	time.AfterFunc(shortContextTimeout, func() {
		cancel()
	})

	return ctx
}

func ContextWithTimeout(timeout time.Duration) context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	time.AfterFunc(timeout, func() {
		cancel()
	})

	return ctx
}
