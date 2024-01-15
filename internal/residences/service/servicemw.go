package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.ResidencesService) domain.ResidencesService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.ResidencesService) domain.ResidencesService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.ResidencesService
}

func (mw loggingMiddleware) CreateResidence(
	ctx context.Context,
	residence *domain.Residence,
	callerID domain.CallerID,
) (result *domain.Residence, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "CreateResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"err", err)
	}()

	return mw.next.CreateResidence(ctx, residence, callerID)
}

func (mw loggingMiddleware) ListResidences(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.Residence, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "ListResidences")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListResidences",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"result", result,
			"err", err,
			"total", total,
		)
	}()

	return mw.next.ListResidences(ctx, criteria, callerID)
}

func (mw loggingMiddleware) GetResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (result *domain.Residence, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "GetResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"err", err)
	}()

	return mw.next.GetResidence(ctx, residenceID, callerID)
}

func (mw loggingMiddleware) UpdateResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	residence *domain.Residence,
	callerID domain.CallerID,
) (result *domain.Residence, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "UpdateResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"err", err)
	}()

	return mw.next.UpdateResidence(ctx, residenceID, residence, callerID)
}

func (mw loggingMiddleware) DeleteResidence(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "DeleteResidence")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteResidence",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"err", err)
	}()

	return mw.next.DeleteResidence(ctx, residenceID, callerID)
}

func (mw loggingMiddleware) IsResidenceExist(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (result bool, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "IsResidenceExist")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "IsResidenceExist",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"residenceID", residenceID,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.IsResidenceExist(ctx, residenceID, callerID)
}

func (mw loggingMiddleware) ListResidencesByIDs(
	ctx context.Context,
	residencesIDs []domain.ResidenceID,
	callerID domain.CallerID,
) (result []*domain.Residence, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "ListResidencesByIDs")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListResidencesByIDs",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"residencesIDs", residencesIDs,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.ListResidencesByIDs(ctx, residencesIDs, callerID)
}

func (mw loggingMiddleware) ListPopularResidences(
	ctx context.Context,
	criteria domain.ResidenceSearchCriteria,
	callerID domain.CallerID,
) (result []*domain.Residence, total domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "ListPopularResidences")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListPopularResidences",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"result", result,
			"criteria", criteria,
			"err", err,
		)
	}()

	return mw.next.ListPopularResidences(ctx, criteria, callerID)
}

func (mw loggingMiddleware) GetConsultationEmailByResidenceID(
	ctx context.Context,
	residenceID domain.ResidenceID,
	callerID domain.CallerID,
) (email string, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "GetConsultationEmailByResidenceID")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetConsultationEmailByResidenceID",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"err", err,
		)
	}()

	return mw.next.GetConsultationEmailByResidenceID(ctx, residenceID, callerID)
}

func (mw loggingMiddleware) CreateFlatPlan(
	ctx context.Context,
	flatPlan *domain.FlatPlan,
	callerID domain.CallerID,
) (result *domain.FlatPlan, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "CreateFlatPlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateFlatPlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"flatPlan", flatPlan,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.CreateFlatPlan(ctx, flatPlan, callerID)
}

func (mw loggingMiddleware) UpdateFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
	flatPlan *domain.FlatPlan,
	callerID domain.CallerID,
) (result *domain.FlatPlan, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "UpdateFlatPlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateFlatPlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"flatPlanID", flatPlanID,
			"flatPlan", flatPlan,
			"result", result,
			"err", err,
		)
	}()

	return mw.next.UpdateFlatPlan(ctx, flatPlanID, flatPlan, callerID)
}

func (mw loggingMiddleware) DeleteFlatPlan(
	ctx context.Context,
	flatPlanID domain.FlatPlanID,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "Residences", "DeleteFlatPlan")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteFlatPlan",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"flatPlanID", flatPlanID,
			"err", err,
		)
	}()

	return mw.next.DeleteFlatPlan(ctx, flatPlanID, callerID)
}
