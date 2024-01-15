package helpers

import (
	"context"

	errors "gitlab.com/zharzhanov/mercury/internal/error"

	"github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/transport"
	"github.com/go-kit/kit/transport/grpc"
	"github.com/go-kit/log"
	"google.golang.org/grpc/metadata"
)

type contextKey string

const (
	identityUserIDKey contextKey = "identity.user_id"
	identityRoleIDKey contextKey = "identity.role_id"
	//identityOrganizationIDKey contextKey = "identity.organization_id"
)

// ServeGrpc - wraps the error
func ServeGrpc(ctx context.Context, req interface{}, handler grpc.Handler) (interface{}, error) {
	_, resp, err := handler.ServeGRPC(ctx, req)
	return resp, errors.GRPCErrorEncoder(err)
}

// SetupServerOptions - setups server options.
func SetupServerOptions(logger log.Logger) []grpc.ServerOption {
	return []grpc.ServerOption{
		grpc.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		grpc.ServerBefore(jwt.GRPCToContext()),
		grpc.ServerBefore(GRPCToContext()),
	}
}

// GRPCToContext moves a JWT from grpc metadata to context. Particularly
// userful for servers.
func GRPCToContext() grpc.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// capital "Key" is illegal in HTTP/2.
		if userIDHeader, ok := md[string(identityUserIDKey)]; ok {
			ctx = context.WithValue(ctx, identityUserIDKey, userIDHeader[0])
		}
		if roleIDHeader, ok := md[string(identityRoleIDKey)]; ok {
			ctx = context.WithValue(ctx, identityRoleIDKey, roleIDHeader[0])
		}
		//if orgIDHeader, ok := md[string(identityOrganizationIDKey)]; ok {
		//	ctx = context.WithValue(ctx, identityOrganizationIDKey, orgIDHeader[0])
		//}
		return ctx
	}
}
