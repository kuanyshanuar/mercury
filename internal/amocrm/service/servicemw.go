package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.AmoCrmService) domain.AmoCrmService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.AmoCrmService) domain.AmoCrmService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.AmoCrmService
}

func (mw loggingMiddleware) CreateApplication(
	ctx context.Context,
	application *domain.AmoApplication,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "AmoCrmService", "CreateApplication")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateApplication",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"application", application,
			"err", err)
	}()
	return mw.next.CreateApplication(ctx, application)
}
