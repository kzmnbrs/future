package future

import (
	"context"
)

// Chain executes futures sequentially, breaking on the first error.
type Chain []Future

func (c *Chain) Await(ctx context.Context) error {
	for _, f := range *c {
		if err := f.Await(ctx); err != nil {
			return err
		}
	}

	return nil
}
