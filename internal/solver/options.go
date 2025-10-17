package solver

import (
	"context"
	"time"
)

// Options configures the solver behavior.
type Options struct {
	MaxSolutions int             // MaxSolutions limits solution search (0 = unlimited)
	Timeout      time.Duration   // Timeout limits solving time
	Randomize    bool            // Randomize solution selection for puzzle generation
	Context      context.Context // Context for cancellation
}

// DefaultOptions returns standard solver options.
func DefaultOptions() *Options {
	return &Options{
		MaxSolutions: 1,
		Randomize:    false,
		Timeout:      1 * time.Second,
	}
}

// GenerateOptions returns solver options useful for puzzle generation.
func GenerateOptions() *Options {
	return &Options{
		MaxSolutions: 1,
		Randomize:    true,
		Timeout:      1 * time.Second,
	}
}

// makeContext creates a context with timeout if specified.
func (s *Solver) makeContext() (context.Context, context.CancelFunc) {
	ctx := s.options.Context
	if ctx == nil {
		ctx = context.Background()
	}

	if s.options.Timeout > 0 {
		return context.WithTimeout(ctx, s.options.Timeout)
	}

	return context.WithCancel(ctx)
}
