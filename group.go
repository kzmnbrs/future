package future

import (
	"context"
	"sync"
)

type Group []Future

func (g *Group) Await(ctx context.Context) error {
	childrenCtx, childrenCancel := context.WithCancel(ctx)
	defer childrenCancel()

	wg := sync.WaitGroup{}
	wg.Add(len(*g))

	var (
		fstErr    error
		fstErrMtx sync.Mutex
	)
	for _, f := range *g {
		go func(f Future) {
			if err := f.Await(childrenCtx); err != nil {
				childrenCancel()

				fstErrMtx.Lock()
				if fstErr == nil {
					fstErr = err
				}
				fstErrMtx.Unlock()
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
		return fstErr
	case <-ctx.Done():
		return ctx.Err()
	}
}
