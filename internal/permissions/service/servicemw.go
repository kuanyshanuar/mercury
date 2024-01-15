package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.PermissionsService) domain.PermissionsService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.PermissionsService) domain.PermissionsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.PermissionsService
}

func (mw loggingMiddleware) CreatePermission(
	ctx context.Context,
	permission *domain.Permission,
	callerID domain.CallerID,
) (result domain.PermissionID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Permissions", "CreatePermission")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreatePermission",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"permission", permission,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.CreatePermission(ctx, permission, callerID)
}

func (mw loggingMiddleware) List(
	ctx context.Context,
	callerID domain.CallerID,
) (result []*domain.Permission, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Permissions", "List")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "List",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.List(ctx, callerID)
}

func (mw loggingMiddleware) Allow(
	ctx context.Context,
	permissionKey string,
	userID domain.UserID,
	roleID domain.RoleID,
	callerID domain.CallerID,
) (result bool, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Permissions", "Allow")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "Allow",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.Allow(ctx, permissionKey, userID, roleID, callerID)
}
