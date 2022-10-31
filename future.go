package future

import (
	"context"
)

type Future interface {
	Await(context.Context) error
}
