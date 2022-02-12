package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/codes"
)

type grpcType string

const (
	// Unary is the grpc method type.
	Unary grpcType = "unary"
)

var (
	// grpcRequestTotal is a metric grpc request counter.
	grpcRequestTotal *prometheus.CounterVec

	// grpcRequestHistogram is the histogram of response latency (seconds) of gRPC.
	grpcRequestHistogram *prometheus.HistogramVec
)

var allCodes = []string{
	codes.OK.String(),
	codes.Canceled.String(),
	codes.Unknown.String(),
	codes.InvalidArgument.String(),
	codes.DeadlineExceeded.String(),
	codes.NotFound.String(),
	codes.AlreadyExists.String(),
	codes.PermissionDenied.String(),
	codes.Unauthenticated.String(),
	codes.ResourceExhausted.String(),
	codes.FailedPrecondition.String(),
	codes.Aborted.String(),
	codes.OutOfRange.String(),
	codes.Unimplemented.String(),
	codes.Internal.String(),
	codes.Unavailable.String(),
	codes.DataLoss.String(),
}

func init() {

	grpcRequestTotal = newCounterVecStartingAtZero(
		prometheus.CounterOpts{
			Name: "grpc_request_total",
			Help: "gRPC request counter.",
		},
		[]string{"grpc_code"}, allCodes...)

	grpcRequestHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_request_duration_milliseconds",
			Help:    "gRPC request histogram.",
			Buckets: prometheus.DefBuckets,
		}, []string{"grpc_type", "grpc_service"})

	initMetrics()
}

func initMetrics() {
	prometheus.MustRegister(
		grpcRequestTotal,
		grpcRequestHistogram,
	)
}

// Reporter represents a metric report.
type Reporter struct {
	rpcType     grpcType
	serviceName string
	methodName  string
	startTime   time.Time
}

// NewReporter creates a reporter to notify metrics.
func NewReporter(rpcType grpcType, serviceName, methodName string) *Reporter {
	return &Reporter{
		rpcType:     rpcType,
		serviceName: serviceName,
		methodName:  methodName,
		startTime:   time.Now(),
	}
}

// GrpcRequestCountInc increment the request conunter by code.
func (r *Reporter) GrpcRequestCountInc(code codes.Code) {
	grpcRequestTotal.WithLabelValues(code.String()).Inc()
}

// GrpcRequestHistogram records the requests for the histogram.
func (r *Reporter) GrpcRequestHistogram() {
	grpcRequestHistogram.
		WithLabelValues(string(r.rpcType), r.serviceName).
		Observe(r.getDuration())
}

func (r *Reporter) getDuration() float64 {
	return time.Since(r.startTime).Seconds()
}

func newCounterVecStartingAtZero(opts prometheus.CounterOpts, labels []string, labelValues ...string) *prometheus.CounterVec {
	counter := prometheus.NewCounterVec(opts, labels)

	for _, label := range labelValues {
		counter.WithLabelValues(label).Add(0)
	}

	return counter
}
