package promise

import "context"

type Func struct {
	Func func() error
}

func FromFunc(f func() error) Promise {
	return &Func{Func: f}
}

func (f *Func) Await(_ context.Context, _ context.CancelFunc) error {
	return f.Func()
}
