package service

import (
	"context"
	"net/http"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(service domain.HTTPClientService) domain.HTTPClientService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.HTTPClientService) domain.HTTPClientService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.HTTPClientService
}

func (mw loggingMiddleware) Get(
	ctx context.Context,
	url string,
) (result *http.Response, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "HTTPClientService", "Get")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "Get",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"url", url,
			"err", err,
		)
	}()
	return mw.next.Get(ctx, url)
}

func (mw loggingMiddleware) Post(
	ctx context.Context,
	url string,
	payload interface{},
) (result *http.Response, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "HTTPClientService", "Post")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "Post",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"url", url,
			"err", err,
		)
	}()
	return mw.next.Post(ctx, url, payload)
}
