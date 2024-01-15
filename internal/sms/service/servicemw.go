package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.SmsService) domain.SmsService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.SmsService) domain.SmsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.SmsService
}

func (mw loggingMiddleware) SendSms(
	ctx context.Context,
	sms *domain.Sms,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "SmsService", "SendSms")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendSms",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"sms", sms,
			"err", err)
	}()

	return mw.next.SendSms(ctx, sms)
}
