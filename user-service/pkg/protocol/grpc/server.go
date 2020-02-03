package grpc

import (
	"context"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"github.com/mesment/microservice/user-service/pkg/api/users/v1"
	"github.com/mesment/microservice/user-service/pkg/logger"
	"github.com/mesment/microservice/user-service/pkg/protocol/grpc/middleware"
)

// RunServer runs gRPC service to publish user service
func RunServer(ctx context.Context, v1API v1.AuthServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+ port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer()
	v1.RegisterAuthServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
