package future

import (
	"context"
	"golang.org/x/sync/errgroup"
)

// Group executes futures concurrently, breaking on the first error.
type Group []Future

func (g *Group) Await(ctx context.Context) error {
	eg, egCtx := errgroup.WithContext(ctx)
	for _, f := range *g {
		eg.Go(func(f Future) func() error {
			return func() error {
				return f.Await(egCtx)
			}
		}(f))
	}
	return eg.Wait()
}
