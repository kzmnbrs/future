package future

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"sync/atomic"
	"testing"
)

func TestGroup_Await(t *testing.T) {
	t.Run("err", func(t *testing.T) {
		v := atomic.Int32{}

		g := Group{
			WrapFunc(func(_ context.Context) error {
				return io.ErrClosedPipe
			}),
		}
		for i := 0; i < 10000; i++ {
			g = append(g, WrapVoidFunc(func(ctx context.Context) {
				if ctx.Err() != nil {
					return
				}
				v.Add(1)
			}))
		}

		assert.Equal(t, io.ErrClosedPipe, g.Await(context.Background()))
		assert.True(t, v.Load() > 0 && v.Load() < 10000)
	})
}
