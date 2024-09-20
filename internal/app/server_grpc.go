package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	desc "github.com/VadimGossip/drs_data_loader/pkg/rate_v1"
)

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	reflection.Register(a.grpcServer)
	desc.RegisterRateV1Server(a.grpcServer, a.serviceProvider.RateImpl(ctx))
	return nil
}

func (a *App) runGRPCServer() error {
	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	logrus.Infof("%s GRPC server is running on: %s", a.name, a.serviceProvider.GRPCConfig().Address())

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}
