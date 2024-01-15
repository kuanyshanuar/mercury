package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.CrmService) domain.CrmService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.CrmService) domain.CrmService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.CrmService
}

// SendResidenceContactDetail implements domain.CrmService.
func (mw loggingMiddleware) SendResidenceContactDetail(
	ctx context.Context,
	content domain.ResidenceContactDetailContent,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CrmService", "SendResidenceContactDetail")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendResidenceContactDetail",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"content", content,
			"err", err,
			"callerID", callerID)
	}()
	return mw.next.SendResidenceContactDetail(ctx, content, callerID)
}
