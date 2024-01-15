package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.ContactDetailsService) domain.ContactDetailsService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.ContactDetailsService) domain.ContactDetailsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.ContactDetailsService
}

func (mw loggingMiddleware) CreateContactDetails(
	ctx context.Context,
	contactDetails *domain.ContactDetails,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "ContactDetails", "CreateContactDetails")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateContactDetails",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"contactDetails", contactDetails,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateContactDetails(ctx, contactDetails, callerID)
}

func (mw loggingMiddleware) CreateResidenceContactDetails(
	ctx context.Context,
	contactDetails *domain.ResidenceContactDetails,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "ContactDetails", "CreateResidenceContactDetails")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateResidenceContactDetails",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"contactDetails", contactDetails,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateResidenceContactDetails(ctx, contactDetails, callerID)
}

func (mw loggingMiddleware) ListContactDetails(
	ctx context.Context,
	callerID domain.CallerID,
) (result []*domain.ContactDetails, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "ContactDetails", "ListContactDetails")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListContactDetails",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.ListContactDetails(ctx, callerID)
}

func (mw loggingMiddleware) ListResidenceContactDetails(
	ctx context.Context,
	criteria domain.ResidenceContactDetailsSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.ResidenceContactDetails, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "ContactDetails", "ListResidenceContactDetails")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListResidenceContactDetails",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"criteria", criteria,
			"total", total,
			"err", err)
	}()

	return mw.next.ListResidenceContactDetails(ctx, criteria, callerID)
}

func (mw loggingMiddleware) MarkAsDelivered(
	ctx context.Context,
	contactID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "ContactDetails", "MarkAsDelivered")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "MarkAsDelivered",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"contactID", contactID,
			"err", err,
			"callerID", callerID,
		)
	}()

	return mw.next.MarkAsDelivered(ctx, contactID, callerID)
}
