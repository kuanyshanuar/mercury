package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.MailSenderService) domain.MailSenderService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.MailSenderService) domain.MailSenderService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.MailSenderService
}

func (mw loggingMiddleware) SendMail(
	ctx context.Context,
	receivers []string,
	content interface{},
	template string,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "MailSenderService", "SendMail")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendMail",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"receivers", receivers,
			"template", template,
			"err", err)
	}()

	return mw.next.SendMail(ctx, receivers, content, template)
}
