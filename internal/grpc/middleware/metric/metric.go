package metric

import (
	"context"
	"strings"

	"github.com/todo-app/internal/metrics"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// UnaryInterceptor returns a new unary server interceptors to report metrics.
func UnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		serviceName, methodName := splitMethodName(info.FullMethod)
		monitor := metrics.NewReporter(metrics.Unary, serviceName, methodName)

		resp, err := handler(ctx, req)
		monitor.GrpcRequestHistogram()
		monitor.GrpcRequestCountInc(status.Code(err))
		return resp, err
	}
}

// return service and method name.
func splitMethodName(fullMethodName string) (string, string) {
	fullMethodName = strings.TrimPrefix(fullMethodName, "/")
	if i := strings.Index(fullMethodName, "/"); i >= 0 {
		return fullMethodName[:i], fullMethodName[i+1:]
	}
	return "unknown", "unknown"
}
