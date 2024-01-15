package service

import (
	"context"
	"github.com/go-kit/log"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.NewsService) domain.NewsService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.NewsService) domain.NewsService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.NewsService
}

func (mw loggingMiddleware) CreateArticle(
	ctx context.Context,
	article *domain.Article,
	callerID domain.CallerID,
) (result int64, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "CreateArticle")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetProfile",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article", article,
			"result", result,
			"err", err)
	}()

	return mw.next.CreateArticle(ctx, article, callerID)
}

func (mw loggingMiddleware) DeleteArticle(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "DeleteArticle")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteArticle",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"err", err)
	}()

	return mw.next.DeleteArticle(ctx, articleID, callerID)
}

func (mw loggingMiddleware) UpdateArticle(
	ctx context.Context,
	articleID int64,
	article *domain.Article,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "UpdateArticle")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UpdateArticle",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"article", article,
			"err", err)
	}()

	return mw.next.UpdateArticle(ctx, articleID, article, callerID)
}

func (mw loggingMiddleware) ListArticles(
	ctx context.Context,
	criteria *domain.NewsSearchCriteria,
	callerID domain.CallerID,
) (articles []*domain.Article, totalArticles domain.Total, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "ListArticles")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ListArticles",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"criteria", criteria,
			"err", err)
	}()

	return mw.next.ListArticles(ctx, criteria, callerID)
}
func (mw loggingMiddleware) GetArticle(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (article *domain.Article, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "GetArticle")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "GetArticle",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"err", err)
	}()

	return mw.next.GetArticle(ctx, articleID, callerID)
}

func (mw loggingMiddleware) AddLike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "AddLike")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "AddLike",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"err", err)
	}()

	return mw.next.AddLike(ctx, articleID, callerID)
}

func (mw loggingMiddleware) DeleteLike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "DeleteLike")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteLike",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"callerID", callerID,
			"err", err)
	}()

	return mw.next.DeleteLike(ctx, articleID, callerID)
}

func (mw loggingMiddleware) AddDislike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "AddDislike")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "AddDislike",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"err", err)
	}()

	return mw.next.AddDislike(ctx, articleID, callerID)
}

func (mw loggingMiddleware) DeleteDislike(
	ctx context.Context,
	articleID int64,
	callerID domain.CallerID,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "NewsService", "DeleteDislike")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "DeleteDislike",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"article_id", articleID,
			"err", err)
	}()

	return mw.next.DeleteDislike(ctx, articleID, callerID)
}
