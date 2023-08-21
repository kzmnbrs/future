package future

import "context"

type (
	_func struct {
		Func func(context.Context) error
	}
	_voidFunc struct {
		Func func(context.Context)
	}
)

// WrapFunc wraps the function call in a future.
// Await returns the resulting error.
func WrapFunc(f func(context.Context) error) Future {
	return &_func{Func: f}
}

// WrapVoidFunc wraps the function call in a future.
// Await returns either nil or context.Cancelled.
func WrapVoidFunc(f func(context.Context)) Future {
	return &_voidFunc{Func: f}
}

func (f *_func) Await(ctx context.Context) error {
	errs := make(chan error)
	go func() {
		errs <- f.Func(ctx)
	}()

	select {
	case err := <-errs:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (f *_voidFunc) Await(ctx context.Context) error {
	done := make(chan struct{})
	go func() {
		defer close(done)
		f.Func(ctx)
	}()

	select {
	case <-done:
	case <-ctx.Done():
	}
	return ctx.Err()
}
