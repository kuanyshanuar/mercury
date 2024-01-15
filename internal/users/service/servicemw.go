package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.UserService) domain.UserService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.UserService) domain.UserService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.UserService
}

func (mw loggingMiddleware) ListUsers(
	ctx context.Context,
	criteria domain.UserSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.User, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Users", "ListUsers")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListUsers",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"total", total,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.ListUsers(ctx, criteria, callerID)
}

func (mw loggingMiddleware) UpdateUser(
	ctx context.Context,
	userID domain.UserID,
	user *domain.User,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Users", "UpdateUser")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateUser",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"user", user,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.UpdateUser(ctx, userID, user, callerID)
}
