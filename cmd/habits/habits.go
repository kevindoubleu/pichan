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
	log := logger.NewLogger("habits.main")
	log.Infow("starting server",
		"url", configs.HABITS_URL,
	)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(logger.NewUnaryServerLogger()),
	)
	pb.RegisterScorecardsServer(grpcServer, service.NewScorecardsServer())

	listener, err := net.Listen("tcp", configs.HABITS_URL)
	if err != nil {
		log.Fatalw("could not listen",
			"url", configs.HABITS_URL,
			"err", err,
		)
	}
	defer listener.Close()

	err = grpcServer.Serve(listener)
	if err != nil {
		log.FatalError("grpc server returned ungracefully", err)
	}
}
