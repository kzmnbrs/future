package promise

import (
	"context"
)

type Future interface {
	Await(context.Context, context.CancelFunc) error
}
