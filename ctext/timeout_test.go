package ctext

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRespectCtx(t *testing.T) {
	sleepyFn := func(ctx context.Context, sleepyTime time.Duration) (time.Duration, error) {
		if sleepyTime < 0 {
			return 0, fmt.Errorf("sleepyTime must be greater than zero")
		}

		time.Sleep(sleepyTime)
		return sleepyTime, nil
	}

	testFn := func(timeout, sleepyTime time.Duration) func(t *testing.T) {
		return func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			dur, err := RespectCtx(ctx, func(ctx context.Context) (time.Duration, error) {
				return sleepyFn(ctx, sleepyTime)
			})

			if timeout <= sleepyTime {
				assert.ErrorIs(t, err, context.DeadlineExceeded)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, dur, time.Duration(0))
				assert.InDelta(t, sleepyTime, dur, float64(100*time.Millisecond))
			}
		}
	}

	t.Run("Finishes before deadline", testFn(3*time.Second, 2*time.Second))
	t.Run("Finishes after deadline", testFn(1*time.Second, 10*time.Second))
}
