// Package flexgrpc provides a default set of configuration for hosting a grpc server in a service.
package flexgrpc

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

// Port is the default gRPC port used in examples.
const Port = 50051

// internal logger for the package.
var logger = log.New(os.Stderr, "flexgrpc: ", 0)

// Config represents configuration for the GRPC server.
type Config struct {
	Addr string
}

// New sets up a new grpc server.
func New(config *Config, options ...grpc.ServerOption) *Server {
	if config == nil {
		config = &Config{}
	}
	if config.Addr == "" {
		var addr string
		if addr = os.Getenv("GRPC_ADDR"); addr == "" {
			addr = net.JoinHostPort("0.0.0.0", strconv.Itoa(Port))
		}
		config.Addr = addr
	}
	return &Server{
		Connection: grpc.NewServer(options...),
		Now:        time.Now,
		addr:       config.Addr,
	}
}

// Server represents a collection of functions for starting and running an RPC server.
type Server struct {
	Connection *grpc.Server
	Now        func() time.Time
	addr       string
}

// Run will start the gRPC server and listen for requests.
func (gs *Server) Run(_ context.Context) error {
	listener, err := net.Listen("tcp", gs.addr)
	if err != nil {
		return err
	}

	if address, ok := listener.Addr().(*net.TCPAddr); ok {
		logger.Printf("serving grpc on http://%s\n", address)
	}

	hsrv := health.NewServer()
	grpc_health_v1.RegisterHealthServer(gs.Connection, hsrv)

	return gs.Connection.Serve(listener)
}

// Halt will attempt to gracefully shut down the server.
func (gs *Server) Halt(_ context.Context) error {
	logger.Printf("shutting down http server...")
	gs.Connection.GracefulStop()
	return nil
}
