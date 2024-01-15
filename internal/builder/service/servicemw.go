package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.BuilderService) domain.BuilderService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.BuilderService) domain.BuilderService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.BuilderService
}

func (mw loggingMiddleware) CreateBuilder(
	ctx context.Context,
	builder *domain.Builder,
	callerID domain.CallerID,
) (result domain.BuilderID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Builders", "CreateBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateBuilder(ctx, builder, callerID)
}

func (mw loggingMiddleware) DeleteBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Builders", "DeleteBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"builderID", builderID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteBuilder(ctx, builderID, callerID)
}

func (mw loggingMiddleware) GetBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	callerID domain.CallerID,
) (result *domain.Builder, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Builders", "GetBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"builderID", builderID,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.GetBuilder(ctx, builderID, callerID)
}

func (mw loggingMiddleware) ListBuilders(
	ctx context.Context,
	criteria domain.BuilderSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.Builder, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Builders", "ListBuilders")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListBuilders",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"callerID", callerID,
			"result", result,
			"total", total,
			"err", err)
	}()

	return mw.next.ListBuilders(ctx, criteria, callerID)
}

func (mw loggingMiddleware) UpdateBuilder(
	ctx context.Context,
	builderID domain.BuilderID,
	builder *domain.Builder,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Builders", "UpdateBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"builderID", builderID,
			"builder", builder,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.UpdateBuilder(ctx, builderID, builder, callerID)
}
