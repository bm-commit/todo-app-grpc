package grpc

import (
	"context"
	"fmt"
	"math"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	gmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1 "github.com/todo-app/internal/grpc/handlers/todoapp/v1"
	"github.com/todo-app/internal/grpc/middleware/logging"
	"github.com/todo-app/internal/grpc/middleware/metric"
	"github.com/todo-app/internal/logger"
	"github.com/todo-app/internal/settings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

// RegisterGrpc represents a grpc service with HTTP endpoints
// through GRPC gateway.
type RegisterGrpc interface {
	RegisterService(g *grpc.Server)
	RegisterGateway(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error
}

// Server is the gRPC server.
type Server struct {
	// Config
	cfg settings.Config

	// Context
	ctx        context.Context
	shutdownFn context.CancelFunc

	// Handlers
	handler *v1.TodoAppServiceHandler

	// Server
	grpc *grpc.Server
}

// NewServer create a gRPC server instance.
func NewServer(cfg settings.Config) (*Server, error) {
	rootCtx, shutdownFn := context.WithCancel(context.Background())

	s := &Server{
		cfg:        cfg,
		ctx:        rootCtx,
		shutdownFn: shutdownFn,
		handler:    v1.NewTodoAppService(),
	}

	s.grpc = s.newGrpcServer()
	return s, nil
}

// Run starts the GRPC server.
func (s *Server) Run() {
	errCh := make(chan error, 1)

	go func() { errCh <- s.runGrpcGateway() }()
	go func() { errCh <- s.runGrpcServer() }()

	go func() {
		for err := range errCh {
			if err != nil {
				logger.Error("failed to run server, %v", err)
				s.Shutdown()
				return
			}
		}
	}()
}

// Shutdown cancel the context and wait for the routines.
func (s *Server) Shutdown() {
	logger.Info("shutdown grpc server started")
	s.shutdownFn()
	logger.Info("shutdown grpc server finished")
}

// GetContext returns the server context.
func (s *Server) GetContext() context.Context {
	return s.ctx
}

func (s *Server) newGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		// UnaryInterceptor
		gmiddleware.WithUnaryServerChain(
			grpcrecovery.UnaryServerInterceptor(),
			logging.UnaryInterceptor(),
			metric.UnaryInterceptor(),
			/*
				auth.UnaryInterceptor(),
				limit.UnaryInterceptor(s.limiterSvc),
			*/
		),

		// Keepalive
		grpc.MaxConcurrentStreams(math.MaxInt32),
	}

	server := grpc.NewServer(opts...)
	reflection.Register(server)
	return server
}

func (s *Server) runGrpcServer() error {
	s.handler.RegisterService(s.grpc)

	go func() {
		<-s.ctx.Done()
		s.grpc.Stop()
		logger.Info("gRPC server on port %d was shutdown", s.cfg.GrpcAPIPort)
	}()

	logger.Info("starting gRPC server on port %d", s.cfg.GrpcAPIPort)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcAPIPort))

	if err != nil {
		return fmt.Errorf("failed to start gRPC server %v", err.Error())
	}
	if err = s.grpc.Serve(listen); err != nil {
		return fmt.Errorf("failed to start gRPC server %v", err.Error())
	}

	return nil
}

func (s *Server) runGrpcGateway() error {
	mux := runtime.NewServeMux(
		runtime.WithErrorHandler(customError),
		runtime.WithOutgoingHeaderMatcher(customOutgoingHeaderMatcher),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions:   protojson.MarshalOptions{EmitUnpopulated: false},
				UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
			},
		}),
	)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	addr := fmt.Sprintf(":%d", s.cfg.GrpcAPIPort)

	err := s.handler.RegisterGateway(s.ctx, mux, addr, opts)
	if err != nil {
		return err
	}

	r := chi.NewRouter()
	r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r)
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.cfg.HTTPAPIPort),
		Handler: r,
	}

	go func() {
		<-s.ctx.Done()
		if err := server.Shutdown(s.ctx); err != nil {
			logger.Error("shutdown gRPC gateway server, %v", err)
			return
		}
		logger.Info("gRPC Gateway on port %d was shutdown", s.cfg.HTTPAPIPort)
	}()

	logger.Info("starting gRPC Gateway on port %d", s.cfg.HTTPAPIPort)

	if err = server.ListenAndServe(); err != http.ErrServerClosed && err != nil {
		logger.Fatal("gateway server %v", err)
	}

	return err
}
