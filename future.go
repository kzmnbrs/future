package future

import (
	"context"
)

//go:generate mockgen -source=$GOFILE -package=future -destination=future_mock.go

type Future interface {
	Await(context.Context) error
}
