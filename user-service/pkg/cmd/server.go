package cmd

import (
	"context"
	"fmt"
	"github.com/mesment/microservice/user-service/pkg/logger"
	"github.com/mesment/microservice/user-service/pkg/protocol/grpc"
	v1 "github.com/mesment/microservice/user-service/pkg/service/v1"
	"github.com/mesment/microservice/user-service/pkg/util"
)


func RunServer(cfg *util.Config) error {
	ctx := context.Background()

	if len(cfg.Server.GrpcPort) == 0 {
		return fmt.Errorf("invalid TCP port for gRPC server: '%s'", cfg.Server.GrpcPort)
	}
	logger.Init(cfg.Log.LogLevel, cfg.Log.LogFile)
	log := logger.NewSugardLogger()

	v1API, err := v1.NewUserService(cfg)
	if err != nil {
		log.Infow("NewUserService", "error", err.Error())
		return  err
	}
	return grpc.RunServer(ctx, v1API, cfg.Server.GrpcPort)
}
