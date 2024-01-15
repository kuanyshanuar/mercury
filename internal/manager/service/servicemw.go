package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.ManagerService) domain.ManagerService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.ManagerService) domain.ManagerService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.ManagerService
}

func (mw loggingMiddleware) CreateManager(
	ctx context.Context,
	manager *domain.Manager,
	callerID domain.CallerID,
) (result domain.ManagerID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Managers", "CreateManager")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateManager",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.CreateManager(ctx, manager, callerID)
}

func (mw loggingMiddleware) GetManager(
	ctx context.Context,
	managerID domain.ManagerID,
	callerID domain.CallerID,
) (result *domain.Manager, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Managers", "GetManager")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetManager",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.GetManager(ctx, managerID, callerID)
}

func (mw loggingMiddleware) ListManagers(
	ctx context.Context,
	criteria domain.ManagerSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.Manager, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Managers", "ListManagers")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListManagers",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.ListManagers(ctx, criteria, callerID)
}

func (mw loggingMiddleware) UpdateManager(
	ctx context.Context,
	managerID domain.ManagerID,
	manager *domain.Manager,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Managers", "UpdateManager")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateManager",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.UpdateManager(ctx, managerID, manager, callerID)
}

func (mw loggingMiddleware) DeleteManager(
	ctx context.Context,
	managerID domain.ManagerID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Managers", "DeleteManager")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteManager",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteManager(ctx, managerID, callerID)
}
