package future

import "context"

type (
	Func struct {
		Func func() error
	}
	VoidFunc struct {
		Func func()
	}
)

func WrapFunc(f func() error) Future {
	return &Func{Func: f}
}

func WrapVoidFunc(f func()) Future {
	return &VoidFunc{Func: f}
}

func (f *Func) Await(ctx context.Context) error {
	res := make(chan error)
	go func() {
		res <- f.Func()
	}()

	select {
	case err := <-res:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (f *VoidFunc) Await(ctx context.Context) error {
	quit := make(chan struct{})
	go func() {
		f.Func()
		quit <- struct{}{}
	}()

	select {
	case <-quit:
	case <-ctx.Done():
	}
	return ctx.Err()
}
