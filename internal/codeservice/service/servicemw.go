package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.CodeService) domain.CodeService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.CodeService) domain.CodeService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.CodeService
}

func (mw loggingMiddleware) SendVerificationCode(
	ctx context.Context,
	phone string,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Identity", "SendVerificationCode")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendVerificationCode",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"phone", phone,
			"err", err,
		)
	}()

	return mw.next.SendVerificationCode(ctx, phone)
}

func (mw loggingMiddleware) ValidateCode(
	ctx context.Context,
	code string,
) (result string, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Identity", "ValidateCode")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ValidateCode",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"code", code,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.ValidateCode(ctx, code)
}

func (mw loggingMiddleware) SendCode(
	ctx context.Context,
	phone string,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Identity", "SendCode")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendCode",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"phone", phone,
			"err", err,
		)
	}()

	return mw.next.SendCode(ctx, phone)
}
