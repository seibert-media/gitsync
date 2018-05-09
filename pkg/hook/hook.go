package hook

import (
	"context"
)

// Hook is responsible for calling a specified endpoint and pass errors through to its caller
type Hook struct {
	URL string
}

// New Hook
func New(url string) *Hook {
	return &Hook{
		URL: url,
	}
}

// Call the specified url
func (h *Hook) Call(ctx context.Context) error {
	return nil
}
