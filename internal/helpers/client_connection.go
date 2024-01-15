package helpers

import (
	"strings"
	"time"

	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// GetGrpcConnection - returns grpc-connection with interceptors.
func GetGrpcConnection(address, secret string) (*grpc.ClientConn, error) {
	var (
		transportCredentials = withTransportCredentials(address)
		secretInterceptor    = NewSecretInterceptor(secret)
	)

	return grpc.Dial(address, transportCredentials,
		grpc.WithChainUnaryInterceptor(
			secretInterceptor.Interceptor, // order matters
		),
		grpc.WithChainStreamInterceptor(
			secretInterceptor.StreamInterceptor, // order matters
		),
		clientKeepAliveParams,
	)
}

func withTransportCredentials(address string) grpc.DialOption {
	if strings.Contains(address, "run.app") {
		// cloud run
		return grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true}))
	}
	return grpc.WithTransportCredentials(insecure.NewCredentials())
}

// ClientKeepAliveParams - gRPC Client Keepalive Parameters
var clientKeepAliveParams = grpc.WithKeepaliveParams(keepalive.ClientParameters{
	// After a duration of this time if the client doesn't see any activity it
	// pings the server to see if the transport is still alive.
	// If set below 10s, a minimum value of 10s will be used instead.
	Time: 5 * time.Second,
	// After having pinged for keepalive check, the client waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout: 5 * time.Second,
	// If true, client sends keepalive pings even with no active RPCs. If false,
	// when there are no active RPCs, Time and Timeout will be ignored and no
	// keepalive pings will be sent.
	PermitWithoutStream: true,
})

// ServerKeepaliveParams - gRPC Server Keepalive Parameters
var ServerKeepaliveParams = grpc.KeepaliveParams(keepalive.ServerParameters{
	// After a duration of this time if the server doesn't see any activity it
	// pings the client to see if the transport is still alive.
	// If set below 1s, a minimum value of 1s will be used instead.
	Time: 5 * time.Second,
	// After having pinged for keepalive check, the server waits for a duration
	// of Timeout and if no activity is seen even after that the connection is
	// closed.
	Timeout: 5 * time.Second,
	// MaxConnectionAgeGrace is an additive period after MaxConnectionAge after
	// which the connection will be forcibly closed.
	MaxConnectionAgeGrace: 10 * time.Second,
})
