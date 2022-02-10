package main

import "flag"

var (
	configFile = flag.String("config", "config/config.yaml", "path to config file")
)

func main() {
	flag.Parse()

	// Load configuration

	// Init storage

	// Init GRPC Server
}
