package logger

import (
	"context"
	"net/http"
)

type loggerKey struct{}

// WithContext returns a new context with the provided logger.
// Use in combination with logger.WithField for great effect.
func WithContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext retrieves the current logger from the context.
func FromContext(ctx context.Context) Logger {
	logger := ctx.Value(loggerKey{})
	if logger == nil {
		return Default
	}
	return logger.(Logger)
}

// FromRequest retrieves the current logger from the request. If no
// logger is available, the default logger is returned.
func FromRequest(r *http.Request) Logger {
	return FromContext(r.Context())
}
