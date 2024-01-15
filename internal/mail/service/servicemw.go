package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.MailService) domain.MailService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.MailService) domain.MailService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.MailService
}

func (mw loggingMiddleware) SendResetPasswordEmail(
	ctx context.Context,
	email string,
	content domain.ResetMailPasswordEmailContent,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "MailService", "SendResetPasswordEmail")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendResetPasswordEmail",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"content", content,
			"err", err)
	}()

	return mw.next.SendResetPasswordEmail(ctx, email, content)
}

func (mw loggingMiddleware) SendResidenceContactDetailEmail(
	ctx context.Context,
	email string,
	content domain.ResidenceContactDetailContent,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "MailService", "SendResidenceContactDetailEmail")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendResidenceContactDetailEmail",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"content", content,
			"err", err)
	}()

	return mw.next.SendResidenceContactDetailEmail(ctx, email, content)
}

func (mw loggingMiddleware) SendContactDetailEmail(
	ctx context.Context,
	email string,
	content domain.ContactDetailContent,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "MailService", "SendContactDetailEmail")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendContactDetailEmail",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"content", content,
			"err", err)
	}()

	return mw.next.SendContactDetailEmail(ctx, email, content)
}
