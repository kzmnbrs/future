package promise

import (
	"context"
)

type Chain []Promise

func (s *Chain) Await(ctx context.Context, cancel context.CancelFunc) error {
	for _, p := range *s {
		if err := p.Await(ctx, cancel); err != nil {
			return err
		}
	}

	return nil
}
