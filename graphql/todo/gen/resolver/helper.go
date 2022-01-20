package resolver

import (
	"context"
	"time"
)

// ResponseInterface 응답
func ResponseInterface(
	ctx context.Context,
	done chan interface{},
	errDone chan interface{},
) (interface{}, error) {
	select {
	case result := <-done:
		return result, nil
	case errResult := <-errDone: // error Code
		return nil, errResult.(error)
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TimeoutContextInit timeout context init
func TimeoutContextInit(
	ctx context.Context,
	timoutDuration time.Duration,
) (
	newContext context.Context,
	doneChan chan interface{},
	errorChan chan interface{},
	cancelFunc context.CancelFunc,
) {
	done := make(chan interface{})
	errDone := make(chan interface{})
	ctx, cancel := context.WithTimeout(ctx, timoutDuration)
	return ctx, done, errDone, cancel
}
