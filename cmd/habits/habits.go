package main

import (
	"net"

	"github.com/kevindoubleu/pichan/configs"
	"github.com/kevindoubleu/pichan/internal/habits/service"
	"github.com/kevindoubleu/pichan/pkg/logger"
	pb "github.com/kevindoubleu/pichan/proto/habits"
	"google.golang.org/grpc"
)

func main() {
	config := loadConfigs()

	log := logger.NewLogger("habits.main")
	log.Infow("starting server",
		"host", config.Server.Host,
		"port", config.Server.Port,
	)

	listener := startTcpListener(log, config)
	closer := getListenerCloser(log, listener)
	defer closer()

	grpcServer := startGrpcServer(config)
	err := grpcServer.Serve(listener)
	if err != nil {
		log.FatalError("grpc server returned ungracefully", err)
	}
}

func loadConfigs() configs.Config {
	config, err := configs.NewConfig(configs.ConfigFile)
	if err != nil {
		panic(err)
	}
	return *config
}

func startTcpListener(log logger.Logger, config configs.Config) net.Listener {
	listener, err := net.Listen("tcp", config.Server.Host+":"+config.Server.Port)
	if err != nil {
		log.Fatalw("could not listen",
			"url", config.Server.Host+":"+config.Server.Port,
			"err", err,
		)
	}
	return listener
}

func getListenerCloser(log logger.Logger, listener net.Listener) func() {
	return func() {
		err := listener.Close()
		if err != nil {
			log.Errorw("could not close listener",
				"listener", listener,
				"err", err,
			)
		}
	}
}

func startGrpcServer(config configs.Config) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.NewUnaryServerLogger()),
	)
	pb.RegisterScorecardsServer(grpcServer, service.NewScorecardsServer(config))
	return grpcServer
}
