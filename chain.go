package promise

import (
	"context"
)

type Chain []Future

func (s *Chain) Await(ctx context.Context, cancel context.CancelFunc) error {
	for _, f := range *s {
		if err := f.Await(ctx, cancel); err != nil {
			return err
		}
	}

	return nil
}
