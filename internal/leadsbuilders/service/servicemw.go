package service

import (
	"context"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.LeadBuilderService) domain.LeadBuilderService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.LeadBuilderService) domain.LeadBuilderService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.LeadBuilderService
}

func (mw loggingMiddleware) CreateLeadBuilder(
	ctx context.Context,
	lead *domain.LeadBuilder,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "CreateLeadBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateLeadBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"err", err)
	}()

	return mw.next.CreateLeadBuilder(ctx, lead, callerID)
}

func (mw loggingMiddleware) GetLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (result *domain.LeadBuilder, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "GetLeadBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetLeadBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"err", err)
	}()

	return mw.next.GetLeadBuilder(ctx, leadID, callerID)
}

func (mw loggingMiddleware) UpdateLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	lead *domain.LeadBuilder,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "UpdateLeadBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateLeadBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"err", err)
	}()

	return mw.next.UpdateLeadBuilder(ctx, leadID, lead, callerID)
}

func (mw loggingMiddleware) ListLeadBuilders(
	ctx context.Context,
	criteria domain.LeadBuilderSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.LeadBuilder, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "ListLeadBuilders")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListLeadBuilders",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"result", result,
			"total", total,
			"err", err,
		)
	}()

	return mw.next.ListLeadBuilders(ctx, criteria, callerID)
}

func (mw loggingMiddleware) DeleteLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "DeleteLeadBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteLeadBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"err", err,
		)
	}()

	return mw.next.DeleteLeadBuilder(ctx, leadID, callerID)
}

func (mw loggingMiddleware) RevokeLeadBuilder(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadBuilders", "RevokeLeadBuilder")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "RevokeLeadBuilder",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"err", err,
		)
	}()

	return mw.next.RevokeLeadBuilder(ctx, leadID, callerID)
}
