package helpers

import (
	"context"
	"gitlab.com/zharzhanov/mercury/internal/domain"
	"strconv"
	"time"

	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
)

// IdentityUserID - extracts the User ID from the context.
func IdentityUserID(ctx context.Context) int64 {
	claims, ok := ctx.Value(jwt.JWTClaimsContextKey).(*Claims)
	if ok {
		return claims.UserID
	}

	return 0
}

// CallerID - extracts the caller ID from the context.
func CallerID(ctx context.Context) (caller domain.CallerID) {
	if val, ok := ctx.Value(identityUserIDKey).(string); ok {
		userID, _ := strconv.Atoi(val)
		caller.UserID = int64(userID)
	}
	if val, ok := ctx.Value(identityRoleIDKey).(string); ok {
		roleID, _ := strconv.Atoi(val)
		caller.RoleID = int64(roleID)
	}
	//if val, ok := ctx.Value(identityOrganizationIDKey).(string); ok {
	//	orgID, _ := strconv.Atoi(val)
	//	caller.OrganizationID = int64(orgID)
	//}
	return
}

// MethodLogger - methods logger.
func MethodLogger(logger log.Logger, s string) endpoint.Middleware {
	return LoggingEndpointMiddleware(log.With(logger, "method", s))
}

// LoggingEndpointMiddleware - logging for mw.
func LoggingEndpointMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				_ = logger.Log("transport_error", err, "took", time.Since(begin))
			}(time.Now())

			return next(ctx, request)
		}
	}
}

// SetupEndpoint - setup endpoint.
func SetupEndpoint(handler endpoint.Endpoint, serviceSecretKey domain.ServiceSecretKey, logger log.Logger, serviceName string, methodName string) endpoint.Endpoint {
	result := createJwtMiddleware(serviceSecretKey)(handler)
	result = MethodLogger(logger, methodName)(result)
	//result = TraceGoKitEndpoint(serviceName, methodName)(result)
	return result
}

// This implementation just skips if the JWTTokenContextKey value is missing
// in the future we might actually want to bail if this happens
func createJwtMiddleware(serviceSecretKey domain.ServiceSecretKey) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// tokenString is stored in the context from the transport handlers.
			tokenString, ok := ctx.Value(jwt.JWTContextKey).(string)

			if !ok {
				return nil, errors.NewErrUnauthorized("missing authentication token")
			}
			if domain.ServiceSecretKey(tokenString) != serviceSecretKey {
				return nil, errors.NewErrUnauthorized("service key is not valid")
			}
			return next(ctx, request) // Allow request if secret token is present
		}
	}
}
