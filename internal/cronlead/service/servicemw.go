package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.CronLeadService) domain.CronLeadService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.CronLeadService) domain.CronLeadService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.CronLeadService
}

func (mw loggingMiddleware) RevokeExpiredLeadResidences(
	ctx context.Context,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CronService", "RevokeExpiredLeadResidences")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "RevokeExpiredLeadResidences",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.RevokeExpiredLeadResidences(ctx, callerID)
}

func (mw loggingMiddleware) RevokeExpiredLeadBuilders(
	ctx context.Context,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CronService", "RevokeExpiredLeadBuilders")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "RevokeExpiredLeadBuilders",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.RevokeExpiredLeadBuilders(ctx, callerID)
}

func (mw loggingMiddleware) RevokeExpiredLeadCottages(
	ctx context.Context,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CronService", "RevokeExpiredLeadCottages")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "RevokeExpiredLeadCottages",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.RevokeExpiredLeadCottages(ctx, callerID)
}
