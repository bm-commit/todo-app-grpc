package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/todo-app/internal/grpc"
	"github.com/todo-app/internal/logger"
	"github.com/todo-app/internal/settings"
)

var (
	configFile = flag.String("config", "config/config.yaml", "path to config file")
)

func main() {
	flag.Parse()

	// Load configuration
	cfg := settings.LoadConfiguration(*configFile)

	// Init storage

	// Init GRPC Server
	srv, err := grpc.NewServer(*cfg)
	if err != nil {
		logger.Fatal("cannot start the server:%v", err)
	}
	srv.Run()
	listenToSystemSignals(srv)
}

func listenToSystemSignals(srv *grpc.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exitChan := make(chan int)

	go func() {
		select {
		case s := <-signalChan:
			switch s {
			case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, os.Kill, os.Interrupt:
				logger.Info("signal received, shutdown server...")
				srv.Shutdown()
				exitChan <- 0
			default:
				logger.Info("unknown sign")
				exitChan <- 1
			}
		case <-srv.GetContext().Done():
			exitChan <- 0
		}
	}()

	code := <-exitChan
	os.Exit(code)
}
