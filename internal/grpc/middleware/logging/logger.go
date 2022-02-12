package logging

import (
	"context"

	"github.com/todo-app/internal/logger"
	"google.golang.org/grpc"
)

// UnaryInterceptor returns a new unary server interceptors to log request.
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			logger.Error("%s. %v", info.FullMethod, err)
		}
		return resp, err
	}
}
