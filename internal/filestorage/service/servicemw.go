package service

import (
	"context"
	"io"

	"gitlab.com/zharzhanov/mercury/internal/domain"

	"github.com/go-kit/log"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.FileStorageService) domain.FileStorageService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.FileStorageService) domain.FileStorageService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.FileStorageService
}

func (mw loggingMiddleware) UploadFile(ctx context.Context, fileInfo *domain.FileInfo, data io.Reader) (result domain.FileURL, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "FileStorage", "UploadFile")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "UploadFile",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"fileInfo", fileInfo,
			"err", err,
			"result", result)
	}()
	return mw.next.UploadFile(ctx, fileInfo, data)
}

func (mw loggingMiddleware) IsFileExist(ctx context.Context, folder string, fileName string) (result bool, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "news", "IsFileExist")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "IsFileExist",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"folder", folder,
			"fileName", fileName,
			"err", err,
			"result", result)
	}()
	return mw.next.IsFileExist(ctx, folder, fileName)
}
