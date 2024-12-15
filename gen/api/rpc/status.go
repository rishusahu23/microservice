package rpc

import "google.golang.org/grpc/codes"

type StatusWithMessage func(msg string) *Status
type StatusWithoutMessage func() *Status

func NewStatusWithMessage(code uint32, msg string) *Status {
	return &Status{
		Code:    code,
		Message: msg,
	}
}

var (
	StatusOk StatusWithoutMessage = func() *Status {
		return NewStatusWithMessage(uint32(codes.OK), "")
	}

	StatusInvalidArgument StatusWithMessage = func(msg string) *Status {
		return NewStatusWithMessage(uint32(codes.InvalidArgument), msg)
	}

	StatusRecordNotFound StatusWithMessage = func(msg string) *Status {
		return NewStatusWithMessage(uint32(codes.NotFound), msg)
	}

	StatusInternal StatusWithMessage = func(msg string) *Status {
		return NewStatusWithMessage(uint32(codes.Internal), msg)
	}
)
