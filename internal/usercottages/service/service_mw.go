package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

type middleware func(userCottage domain.UserCottageService) domain.UserCottageService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.UserCottageService) domain.UserCottageService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.UserCottageService
}

func (mw loggingMiddleware) AddCottageToFavourites(ctx context.Context, userID int64, CottageID int64, callerID domain.CallerID) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserCottage", "AddCottageToFavourites")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "AddCottageToFavourites",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"cottageID", CottageID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.AddCottageToFavourites(ctx, userID, CottageID, callerID)
}

func (mw loggingMiddleware) DeleteCottageFromFavourites(ctx context.Context, userID int64, CottageID int64, callerID domain.CallerID) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserCottage", "DeleteCottageFromFavourites")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteCottageFromFavourites",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"cottageID", CottageID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteCottageFromFavourites(ctx, userID, CottageID, callerID)
}

func (mw loggingMiddleware) ListFavouriteCottages(ctx context.Context, userID int64, criteria domain.FavouriteCottagesSearchCriteria, callerID domain.CallerID) (result []int64, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserCottage", "ListFavouriteCottages")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListFavouriteCottages",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"criteria", criteria,
			"callerID", callerID,
			"result", result,
			"total", total,
			"err", err)
	}()

	return mw.next.ListFavouriteCottages(ctx, userID, criteria, callerID)
}
