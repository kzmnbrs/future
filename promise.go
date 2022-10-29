package promise

import (
	"context"
)

type Promise interface {
	Await(context.Context, context.CancelFunc) error
}
