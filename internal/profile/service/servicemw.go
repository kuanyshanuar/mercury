package service

import (
	"context"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.ProfileService) domain.ProfileService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.ProfileService) domain.ProfileService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.ProfileService
}

func (mw loggingMiddleware) GetProfile(
	ctx context.Context,
	userID domain.UserID,
	callerID domain.CallerID,
) (result *domain.Profile, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "GetProfile")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetProfile",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"result", result,
			"err", err)
	}()

	return mw.next.GetProfile(ctx, userID, callerID)
}

func (mw loggingMiddleware) UpdateProfile(
	ctx context.Context,
	userID domain.UserID,
	profile *domain.Profile,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "UpdateProfile")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateProfile",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"profile", profile,
			"err", err)
	}()

	return mw.next.UpdateProfile(ctx, userID, profile, callerID)
}

func (mw loggingMiddleware) ValidatePhone(
	ctx context.Context,
	code string,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "ValidatePhone")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ValidatePhone",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"code", code,
			"err", err)
	}()

	return mw.next.ValidatePhone(ctx, code, callerID)
}
