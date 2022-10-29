package future

import (
	"context"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Group []Future

func (s *Group) Await(ctx context.Context, cancel context.CancelFunc) error {
	wg := sync.WaitGroup{}
	wg.Add(len(*s))

	var fstErrPtr unsafe.Pointer
	for _, f := range *s {
		go func(f Future) {
			if err := f.Await(ctx, cancel); err != nil {
				atomic.CompareAndSwapPointer(&fstErrPtr, nil, unsafe.Pointer(&err))
			}
			wg.Done()
		}(f)
	}

	quit := make(chan struct{})
	go func() {
		wg.Wait()
		quit <- struct{}{}
	}()

	select {
	case <-quit:
	case <-ctx.Done():
	}

	if errPtr := atomic.LoadPointer(&fstErrPtr); errPtr != nil {
		cancel()
		return *(*error)(errPtr)
	}
	return ctx.Err()
}
