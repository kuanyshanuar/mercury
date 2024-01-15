package service

import (
	"context"

	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.FiltersService) domain.FiltersService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.FiltersService) domain.FiltersService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.FiltersService
}

func (mw loggingMiddleware) ListCities(
	ctx context.Context,
	criteria domain.CitySearchCriteria,
	callerID domain.CallerID,
) (result []*domain.City, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "ListCities")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListCities",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"result", result,
			"total", total,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.ListCities(ctx, criteria, callerID)
}

func (mw loggingMiddleware) ListDistricts(
	ctx context.Context,
	criteria domain.DistrictSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.District, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "ListDistricts")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListDistricts",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"callerID", callerID,
			"result", result,
			"total", total,
			"err", err)
	}()

	return mw.next.ListDistricts(ctx, criteria, callerID)
}

func (mw loggingMiddleware) ListFilters(
	ctx context.Context,
	callerID domain.CallerID,
) (result map[string][]*domain.Filter, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "ListFilters")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListFilters",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.ListFilters(ctx, callerID)
}

func (mw loggingMiddleware) ListBuilders(
	ctx context.Context,
	callerID domain.CallerID,
) (result []*domain.FilterBuilder, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "ListBuilders")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListBuilders",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"err", err)
	}()

	return mw.next.ListBuilders(ctx, callerID)
}

func (mw loggingMiddleware) CreateFilter(
	ctx context.Context,
	key string,
	filter *domain.Filter,
	callerID domain.CallerID,
) (result *domain.Filter, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "CreateFilter")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateFilter",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"key", key,
			"filter", filter,
			"result", result,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.CreateFilter(ctx, key, filter, callerID)
}

func (mw loggingMiddleware) DeleteFilter(
	ctx context.Context,
	id int64,
	key string,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "DeleteFilter")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteFilter",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"id", id,
			"key", key,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.DeleteFilter(ctx, id, key, callerID)
}

func (mw loggingMiddleware) CreateCity(
	ctx context.Context,
	city *domain.City,
	callerID domain.CallerID,
) (result domain.CityID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "CreateCity")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateCity",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"city", city,
			"callerID", callerID,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.CreateCity(ctx, city, callerID)
}

func (mw loggingMiddleware) UpdateCity(
	ctx context.Context,
	cityID domain.CityID,
	city *domain.City,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "UpdateCity")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateCity",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"city", city,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.UpdateCity(ctx, cityID, city, callerID)
}

func (mw loggingMiddleware) DeleteCity(
	ctx context.Context,
	cityID domain.CityID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "DeleteCity")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteCity",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.DeleteCity(ctx, cityID, callerID)
}

func (mw loggingMiddleware) CreateDistrict(
	ctx context.Context,
	district *domain.District,
	callerID domain.CallerID,
) (result domain.DistrictID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "CreateDistrict")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateDistrict",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.CreateDistrict(ctx, district, callerID)
}

func (mw loggingMiddleware) UpdateDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
	district *domain.District,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "UpdateDistrict")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateDistrict",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.UpdateDistrict(ctx, districtID, district, callerID)
}

func (mw loggingMiddleware) DeleteDistrict(
	ctx context.Context,
	districtID domain.DistrictID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Filters", "DeleteDistrict")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteDistrict",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"callerID", callerID,
			"err", err,
		)
	}()

	return mw.next.DeleteDistrict(ctx, districtID, callerID)
}
