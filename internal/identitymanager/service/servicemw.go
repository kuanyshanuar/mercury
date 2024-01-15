package service

import (
	"context"

	"gitlab.com/zharzhanov/mercury/internal/domain"
	telemetry "gitlab.com/zharzhanov/mercury/pkg/opentelemetry"

	"github.com/go-kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type middleware func(newsService domain.IdentityManagerService) domain.IdentityManagerService

// LoggingServiceMiddleware takes a logger as a dependency
// and returns a service Middleware.
func loggingServiceMiddleware(logger log.Logger) middleware {
	return func(next domain.IdentityManagerService) domain.IdentityManagerService {
		return loggingMiddleware{logger, next}
	}
}

type loggingMiddleware struct {
	logger log.Logger
	next   domain.IdentityManagerService
}

func (mw loggingMiddleware) CreateUser(
	ctx context.Context,
	user *domain.User,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "CreateUser")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "CreateUser",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"user", user,
			"err", err)
	}()

	return mw.next.CreateUser(ctx, user)
}

func (mw loggingMiddleware) ValidateUser(
	ctx context.Context,
	email string,
	password string,
) (result *domain.User, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "ValidateUser")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ValidateUser",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"password", password,
			"err", err,
			"result", result,
		)
	}()

	return mw.next.ValidateUser(ctx, email, password)
}

func (mw loggingMiddleware) ValidateCode(
	ctx context.Context,
	code string,
) (userID domain.UserID, roleID domain.RoleID, err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "ValidateCode")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ValidateCode",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"code", code,
			"userID", userID,
			"roleID", roleID,
			"err", err,
		)
	}()

	return mw.next.ValidateCode(ctx, code)
}

func (mw loggingMiddleware) SendResetPasswordToken(
	ctx context.Context,
	email string,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "SendPasswordResetToken")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "SendPasswordResetToken",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"email", email,
			"err", err,
		)
	}()

	return mw.next.SendResetPasswordToken(ctx, email)
}

func (mw loggingMiddleware) ResetPassword(
	ctx context.Context,
	token string,
	newPassword string,
) (err error) {
	ctx, span := telemetry.StartServiceSpan(ctx, "IdentityService", "ResetPassword")
	defer func() {
		traceID, spanID := telemetry.EndSpan(span, err)
		_ = mw.logger.Log("method", "ResetPassword",
			domain.LogFieldTraceID, traceID,
			domain.LogFieldSpanID, spanID,
			"token", token,
			"newPassword", newPassword,
			"err", err,
		)
	}()

	return mw.next.ResetPassword(ctx, token, newPassword)
}
