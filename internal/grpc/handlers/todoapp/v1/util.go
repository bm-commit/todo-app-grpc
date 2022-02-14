package v1

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	return status.Error(codes.Unknown, "Cannot obtain the resource")
}
