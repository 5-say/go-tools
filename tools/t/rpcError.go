package t

import (
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// CustomRPCError ..
type CustomRPCError struct {
	Code           codes.Code
	PrivateMessage string
	PublicMessage  string
}

// RPCError .. 构造 rpc error
func RPCError(privateMessage, publicMessage string) error {
	return grpc.Errorf(codes.InvalidArgument, privateMessage+"{[']}"+publicMessage)
}

// RPCErrorCode .. 构造 rpc error
func RPCErrorCode(privateMessage, publicMessage string, code codes.Code) error {
	return grpc.Errorf(code, privateMessage+"{[']}"+publicMessage)
}

// RPCErrorParse .. 解析 rpc error
func RPCErrorParse(err error) CustomRPCError {
	errorDesc := strings.Split(grpc.ErrorDesc(err), "{[']}")
	return CustomRPCError{
		Code:           grpc.Code(err),
		PrivateMessage: errorDesc[0],
		PublicMessage:  errorDesc[1],
	}
}
