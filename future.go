package future

import (
	"context"
)

// Future represents a deferred function call.
// Wrapped function starts on Await.
type Future interface {
	// Await waits for the future to complete and returns the result.
	Await(context.Context) error
}
