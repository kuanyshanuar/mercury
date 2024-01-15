package error

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Mapping between grpc and http codeservice.
var grpcHTTPCodes = map[int]int{
	0:  200, //"OK",
	1:  499, //"CANCELLED",
	2:  500, //"UNKNOWN",
	3:  400, //"INVALID_ARGUMENT",
	4:  504, //"DEADLINE_EXCEEDED",
	5:  404, //"NOT_FOUND",
	6:  409, //"ALREADY_EXISTS",
	7:  403, //"PERMISSION_DENIED",
	8:  429, //"RESOURCE_EXHAUSTED",
	9:  400, //"FAILED_PRECONDITION",
	10: 409, //"ABORTED",
	11: 400, //"OUT_OF_RANGE",
	12: 501, //"UNIMPLEMENTED",
	13: 500, //"INTERNAL",
	14: 503, //"UNAVAILABLE",
	15: 500, //"DATA_LOSS",
	16: 401, //"UNAUTHENTICATED",
}

// HTTPCodeForGrpc - returns mapped http code for the provided grpc code.
func HTTPCodeForGrpc(c codes.Code) int {
	return grpcHTTPCodes[int(c)]
}

// GRPCErrorEncoder - returns mapped grpc error based on the provided error.
func GRPCErrorEncoder(err error) error {
	if err == nil {
		return err
	}
	switch err.(type) {
	case *ErrInvalidArgument:
		return status.Error(codes.InvalidArgument, err.Error())
	case *ErrAlreadyExist:
		return status.Error(codes.AlreadyExists, err.Error())
	case *ErrNotFound:
		return status.Error(codes.NotFound, err.Error())
	case *ErrFailedPrecondition:
		return status.Error(codes.FailedPrecondition, err.Error())
	case *ErrInternal:
		return status.Error(codes.Internal, err.Error())
	case *ErrUnauthorized:
		return status.Error(codes.Unauthenticated, err.Error())
	case *ErrPermissionDenied:
		return status.Error(codes.PermissionDenied, err.Error())
	default:
		return status.Error(codes.Unknown, err.Error())
	}
}
