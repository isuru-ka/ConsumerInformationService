package middleware

import (
	"consumerinformationmodule/pkg/env"
	"google.golang.org/grpc"
)

func GetGrpcServerConfigs() *grpc.Server{
	s := grpc.NewServer(
	)
	return s
}

func GetGrpcPort() string {
    return env.GRPCPORT
}