package jagwerror

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorCode string

const (
	UNKNOWN           ErrorCode = "UNKNOWN"
	INVALID_ARGUMENT  ErrorCode = "INVALID_ARGUMENT"
	DEADLINE_EXCEEDED ErrorCode = "DEADLINE_EXCEEDED"
	NOT_FOUND         ErrorCode = "NOT_FOUND"
	UNIMPLEMENTED     ErrorCode = "UNIMPLEMENTED"
	INTERNAL          ErrorCode = "INTERNAL"
	UNAVAILABLE       ErrorCode = "UNAVAILABLE"
)

type Error struct {
	ErrorCode ErrorCode
	Message   string
}

func CreateErrorForKeysNotFound(keys []string) *Error {
	if len(keys) == 0 {
		return nil
	}

	keysString := strings.Join(keys, ", ")
	message := "Unable to find the following keys: " + keysString

	return &Error{
		ErrorCode: NOT_FOUND,
		Message: message,
	}
}

func GetGrpcError(err *Error) error {
	return status.Errorf(GetGrpcErrorCode(err), err.Message)
}

func GetGrpcErrorCode(err *Error) codes.Code {
	switch err.ErrorCode {
	case "INVALID_ARGUMENT":
		return codes.InvalidArgument
	case "DEADLINE_EXCEEDED":
		return codes.DeadlineExceeded
	case "NOT_FOUND":
		return codes.NotFound
	case "UNIMPLEMENTED":
		return codes.Unimplemented
	case "INTERNAL":
		return codes.Internal
	case "UNAVAILABLE":
		return codes.Unavailable
	default:
		return codes.Unknown
	}
}