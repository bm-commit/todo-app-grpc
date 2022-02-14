package grpc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/todo-app/internal/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type errorBody struct {
	Code   uint32        `json:"code,omitempty"`
	Err    string        `json:"error,omitempty"`
	Detail []interface{} `json:"detail,omitempty"`
}

func customOutgoingHeaderMatcher(key string) (string, bool) {
	switch key {
	case
		"x-ratelimit-limit",
		"x-ratelimit-remaining",
		"x-ratelimit-reset":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

func customError(_ context.Context, _ *runtime.ServeMux, marshaller runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Set("Content-type", marshaller.ContentType(s.Proto()))
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))

	jErr := json.NewEncoder(w).Encode(errorBody{
		Code:   uint32(s.Code()),
		Err:    s.Message(),
		Detail: s.Details(),
	})

	if jErr != nil {
		if _, err := w.Write([]byte(fallback)); err != nil {
			logger.Error("failed to write the response: %v", err)
		}
	}
}
