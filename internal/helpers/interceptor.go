package helpers

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// SecretInterceptor - use secret key for microservices communication outside VPC.
type SecretInterceptor struct {
	secret string
}

// NewSecretInterceptor - just constructor
func NewSecretInterceptor(secret string) *SecretInterceptor {
	return &SecretInterceptor{secret: secret}
}

// Interceptor - use secret key for microservices communication outside VPC.
func (i *SecretInterceptor) Interceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", i.secret)

	return invoker(ctx, method, req, reply, cc, opts...)
}

// StreamInterceptor - use secret key for microservices communication outside VPC.
func (i *SecretInterceptor) StreamInterceptor(
	ctx context.Context,
	desc *grpc.StreamDesc,
	cc *grpc.ClientConn,
	method string,
	streamer grpc.Streamer,
	opts ...grpc.CallOption,
) (grpc.ClientStream, error) {
	ctx = metadata.AppendToOutgoingContext(ctx, "authorization", i.secret)

	return streamer(ctx, desc, cc, method, opts...)
}
