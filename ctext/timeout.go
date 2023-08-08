package ctext

import "context"

// CtxFn is a function that takes a context and returns a value and an error
type CtxFn[T any] func(ctx context.Context) (T, error)

// RespectCtx runs a function with a context and returns the result or an error
func RespectCtx[T any](ctx context.Context, fn CtxFn[T]) (T, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	resChan := make(chan T)
	errChan := make(chan error)

	wf := func(c chan<- T, e chan<- error) {
		res, err := fn(ctx)

		if err != nil {
			errChan <- err
			return
		}

		resChan <- res
	}

	go wf(resChan, errChan)
	var empty T

	select {
	case <-ctx.Done():
		return empty, ctx.Err()
	case err := <-errChan:
		return empty, err
	case res := <-resChan:
		return res, nil
	}
}
