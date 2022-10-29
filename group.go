package promise

import (
	"context"
	"sync"
	"sync/atomic"
	"unsafe"
)

type Group []Promise

func (s *Group) Await(ctx context.Context, cancel context.CancelFunc) error {
	wg := sync.WaitGroup{}
	wg.Add(len(*s))

	var fstErrPtr unsafe.Pointer
	for _, p := range *s {
		go func(p Promise) {
			if err := p.Await(ctx, cancel); err != nil {
				atomic.StorePointer(&fstErrPtr, unsafe.Pointer(&err))
				cancel()
			}
			wg.Done()
		}(p)
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
		return *(*error)(errPtr)
	}
	return ctx.Err()
}
