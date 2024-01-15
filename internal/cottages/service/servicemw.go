package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(cottageService domain.CottageService) domain.CottageService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.CottageService) domain.CottageService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.CottageService
}

func (mw loggingMiddleware) UpdateHousePlan(
	ctx context.Context,
	housePlanID int64,
	housePlan *domain.HousePlan,
	callerID domain.CallerID,
) (result *domain.HousePlan, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "UpdateHousePlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateHousePlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"housePlanID", housePlanID,
			"housePlan", housePlan,
			"result", result,
			"err", err)
	}()

	return mw.next.UpdateHousePlan(ctx, housePlanID, housePlan, callerID)
}

func (mw loggingMiddleware) DeleteHousePlan(
	ctx context.Context,
	housePlanID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "DeleteHousePlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteHousePlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"housePlanID", housePlanID,
			"err", err)
	}()

	return mw.next.DeleteHousePlan(ctx, housePlanID, callerID)
}

func (mw loggingMiddleware) CreateHousePlan(
	ctx context.Context,
	housePlan *domain.HousePlan,
	callerID domain.CallerID,
) (result *domain.HousePlan, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "CreateHousePlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateHousePlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"housePlan", housePlan,
			"result", result,
			"err", err)
	}()

	return mw.next.CreateHousePlan(ctx, housePlan, callerID)
}

func (mw loggingMiddleware) UpdateCottage(
	ctx context.Context,
	cottageID int64,
	cottage *domain.Cottage,
	callerID domain.CallerID,
) (cottageResult *domain.Cottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "UpdateCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottageID", cottageID,
			"updated", cottageResult,
			"cottage", cottage,
			"err", err)
	}()

	return mw.next.UpdateCottage(ctx, cottageID, cottage, callerID)
}

func (mw loggingMiddleware) DeleteCottage(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "DeleteCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottageID", cottageID,
			"err", err)
	}()

	return mw.next.DeleteCottage(ctx, cottageID, callerID)

}

func (mw loggingMiddleware) CreateCottage(
	ctx context.Context,
	cottage *domain.Cottage,
	callerID domain.CallerID,
) (cottageResult *domain.Cottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "CreateCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottage", cottage,
			"cottage_status", cottage.StatusID,
			"cottage_city", cottage.CityID,
			"cottage_result", cottageResult,
			"err", err)
	}()

	return mw.next.CreateCottage(ctx, cottage, callerID)
}

func (mw loggingMiddleware) GetCottage(
	ctx context.Context,
	id int64,
	callerID domain.CallerID,
) (cottage *domain.Cottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "GetCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"id", id,
			"cottage", cottage,
			"err", err)
	}()

	return mw.next.GetCottage(ctx, id, callerID)
}

func (mw loggingMiddleware) ListCottage(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
	callerID domain.CallerID,
) (cottages []*domain.Cottage, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "ListCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"cottages", cottages,
			"total", total,
			"err", err)
	}()

	return mw.next.ListCottage(ctx, criteria, callerID)
}

func (mw loggingMiddleware) ListPopularCottages(
	ctx context.Context,
	criteria domain.CottageSearchCriteria,
	callerID domain.CallerID,
) (cottages []*domain.Cottage, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "ListPopularCottages")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListPopularCottages",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"cottages", cottages,
			"total", total,
			"err", err)
	}()

	return mw.next.ListPopularCottages(ctx, criteria, callerID)
}

func (mw loggingMiddleware) ListCottagesByIDs(
	ctx context.Context,
	id []int64,
	callerID domain.CallerID,
) (cottages []*domain.Cottage, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "ListCottagesById")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListCottagesById",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"ids", id,
			"cottages", cottages,
			"err", err)
	}()

	return mw.next.ListCottagesByIDs(ctx, id, callerID)
}

func (mw loggingMiddleware) IsFavouriteCottage(
	ctx context.Context,
	cottageID int64,
	userID int64,
	callerID domain.CallerID,
) (res bool, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "IsFavouriteCottage")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "IsFavouriteCottage",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottageID", cottageID,
			"userID", userID,
			"res", res,
			"err", err)
	}()

	return mw.next.IsFavouriteCottage(ctx, cottageID, userID, callerID)
}

func (mw loggingMiddleware) IsCottageExist(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) (res bool, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "IsCottageExist")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "IsCottageExist",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottageID", cottageID,
			"res", res,
			"err", err)
	}()

	return mw.next.IsCottageExist(ctx, cottageID, callerID)
}

func (mw loggingMiddleware) GetConsultationEmailByCottageID(
	ctx context.Context,
	cottageID int64,
	callerID domain.CallerID,
) (res string, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "CottageService", "GetConsultationEmailByCottageID")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetConsultationEmailByCottageID",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"cottageID", cottageID,
			"result", res,
			"err", err)
	}()

	return mw.next.GetConsultationEmailByCottageID(ctx, cottageID, callerID)
}
