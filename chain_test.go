package future

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestChain_Await(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		i := 0
		c := Chain{
			WrapVoidFunc(func(_ context.Context) {
				i++
			}),
			WrapVoidFunc(func(_ context.Context) {
				i++
			}),
		}

		assert.Nil(t, c.Await(context.Background()))
		assert.Equal(t, 2, i)
	})

	t.Run("err", func(t *testing.T) {
		i := 0
		c := Chain{
			WrapVoidFunc(func(_ context.Context) {
				i++
			}),
			WrapVoidFunc(func(_ context.Context) {
				i++
			}),
			WrapFunc(func(_ context.Context) error {
				return io.ErrClosedPipe
			}),
		}

		err := c.Await(context.Background())
		assert.Equal(t, 2, i)
		assert.Equal(t, io.ErrClosedPipe, err)
	})

	t.Run("ctx cancel", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		c := Chain{
			WrapFunc(func(_ context.Context) error {
				time.Sleep(time.Second)
				return nil
			}),
			WrapVoidFunc(func(_ context.Context) {}),
		}
		assert.Equal(t, context.Canceled, c.Await(ctx))
	})
}
