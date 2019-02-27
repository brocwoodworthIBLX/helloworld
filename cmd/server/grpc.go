package main

import (
	"github.com/bwoodworthIBLX/helloworld/pkg/pb"
	"github.com/bwoodworthIBLX/helloworld/pkg/svc"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewGRPCServer(logger *logrus.Logger) (*grpc.Server, error) {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				// Request-Id interceptor
				requestid.UnaryServerInterceptor(),

				// logging middleware
				grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(logger)),

				// validation middleware
				grpc_validator.UnaryServerInterceptor(),

				// collection operators middleware
				gateway.UnaryServerInterceptor(),
			),
		),
	)

	// register service implementation with the grpcServer
	s, err := svc.NewBasicServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterHelloworldServer(grpcServer, s)

	return grpcServer, nil
}
