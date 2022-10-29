package future

import "context"

type Func struct {
	Func func() error
}

func FromFunc(f func() error) Future {
	return &Func{Func: f}
}

func (f *Func) Await(ctx context.Context, cancel context.CancelFunc) error {
	res := make(chan error)
	go func() {
		res <- f.Func()
	}()

	select {
	case err := <-res:
		if err != nil {
			cancel()
		}
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
