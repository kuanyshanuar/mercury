package service

import (
	"context"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"gitlab.com/zharzhanov/mercury/internal/helpers"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
	"time"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.LeadCottageService) domain.LeadCottageService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.LeadCottageService) domain.LeadCottageService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.LeadCottageService
}

func (mw loggingMiddleware) CreateLeadCottage(ctx context.Context, lead *domain.LeadCottage, callerID domain.CallerID) (id int64, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "leadCottages", "CreateLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"lead", lead,
			"current_date:", helpers.EndOfTheDay(time.Now().Unix()),
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateLeadCottage(ctx, lead, callerID)
}

func (mw loggingMiddleware) UpdateLeadCottage(ctx context.Context, leadID int64, lead *domain.LeadCottage, callerID domain.CallerID) (result *domain.LeadCottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadCottages", "UpdateLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"lead", lead,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.UpdateLeadCottage(ctx, leadID, lead, callerID)
}

func (mw loggingMiddleware) DeleteLeadCottage(ctx context.Context, leadID int64, callerID domain.CallerID) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadCottages", "DeleteLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteLeadCottage(ctx, leadID, callerID)
}

func (mw loggingMiddleware) GetLeadCottage(ctx context.Context, leadID int64, callerID domain.CallerID) (result *domain.LeadCottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadCottages", "GetLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"callerID", callerID,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.GetLeadCottage(ctx, leadID, callerID)
}

func (mw loggingMiddleware) ListLeadCottage(ctx context.Context, criteria domain.LeadCottageSearchCriteria, callerID domain.CallerID) (result []*domain.LeadCottage, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadCottages", "ListLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"callerID", callerID,
			"result", result,
			"total", total,
			"err", err,
		)
	}()

	return mw.next.ListLeadCottage(ctx, criteria, callerID)
}
func (mw loggingMiddleware) RevokeLeadCottage(ctx context.Context, leadID domain.LeadID, callerID domain.CallerID) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "LeadCottages", "ListLeadCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListLeadCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"leadID", leadID,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.RevokeLeadCottage(ctx, leadID, callerID)
}
