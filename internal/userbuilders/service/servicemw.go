package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.UserBuilderService) domain.UserBuilderService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.UserBuilderService) domain.UserBuilderService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.UserBuilderService
}

func (mw loggingMiddleware) Subscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserBuilderService", "Subscribe")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "Subscribe",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"subscriberID", subscriberID,
			"builderID", builderID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.Subscribe(ctx, subscriberID, builderID, callerID)
}

func (mw loggingMiddleware) Unsubscribe(
	ctx context.Context,
	subscriberID int64,
	builderID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserBuilderService", "Unsubscribe")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "Unsubscribe",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"subscriberID", subscriberID,
			"builderID", builderID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.Unsubscribe(ctx, subscriberID, builderID, callerID)
}

func (mw loggingMiddleware) ListBuilders(
	ctx context.Context,
	subscriberID int64,
	criteria domain.UserBuilderSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.Builder, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserBuilderService", "ListBuilders")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListBuilders",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"subscriberID", subscriberID,
			"criteria", criteria,
			"result", result,
			"total", total,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.ListBuilders(ctx, subscriberID, criteria, callerID)
}
