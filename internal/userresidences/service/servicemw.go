package service

import (
	"context"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.UserResidenceService) domain.UserResidenceService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.UserResidenceService) domain.UserResidenceService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.UserResidenceService
}

func (mw loggingMiddleware) AddResidenceToFavourites(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserResidence", "AddResidenceToFavourites")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "AddResidenceToFavourites",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"residenceID", residenceID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.AddResidenceToFavourites(ctx, userID, residenceID, callerID)
}

func (mw loggingMiddleware) DeleteResidenceFromFavourites(
	ctx context.Context,
	userID domain.UserID,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserResidence", "DeleteResidenceFromFavourites")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteResidenceFromFavourites",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"residenceID", residenceID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteResidenceFromFavourites(ctx, userID, residenceID, callerID)
}

func (mw loggingMiddleware) ListFavouriteResidences(
	ctx context.Context,
	userID domain.UserID,
	criteria domain.FavouriteResidencesSearchCriteria,
	callerID domain.CallerID,
) (result []int64, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "UserResidence", "ListFavouriteResidences")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListFavouriteResidences",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"userID", userID,
			"criteria", criteria,
			"result", result,
			"total", total,
			"err", err)
	}()

	return mw.next.ListFavouriteResidences(ctx, userID, criteria, callerID)
}
