package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.LeadResidenceService) domain.LeadResidenceService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.LeadResidenceService) domain.LeadResidenceService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.LeadResidenceService
}

func (mw loggingMiddleware) CreateLeadResidence(
	ctx context.Context,
	lead *domain.LeadResidence,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "CreateLeadResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateLeadResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"lead", lead,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateLeadResidence(ctx, lead, callerID)
}

func (mw loggingMiddleware) ListLeadResidences(
	ctx context.Context,
	criteria domain.LeadResidenceSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.LeadResidence, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "ListLeadResidences")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListLeadResidences",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"callerID", callerID,
			"result", result,
			"total", total,
			"err", err,
		)
	}()

	return mw.next.ListLeadResidences(ctx, criteria, callerID)
}

func (mw loggingMiddleware) UpdateLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	lead *domain.LeadResidence,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "UpdateLeadResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateLeadResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"lead", lead,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.UpdateLeadResidence(ctx, leadID, lead, callerID)
}

func (mw loggingMiddleware) DeleteLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "DeleteLeadResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteLeadResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteLeadResidence(ctx, leadID, callerID)
}

func (mw loggingMiddleware) RevokeLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "RevokeLeadResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "RevokeLeadResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.RevokeLeadResidence(ctx, leadID, callerID)
}

func (mw loggingMiddleware) GetLeadResidence(
	ctx context.Context,
	leadID domain.LeadID,
	callerID domain.CallerID,
) (result *domain.LeadResidence, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadResidences", "GetLeadResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetLeadResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"result", result,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.GetLeadResidence(ctx, leadID, callerID)
}
