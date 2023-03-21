package logger

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
)

func NewUnaryServerLogger() grpc.UnaryServerInterceptor {
	return grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(NewLogger("habits.grpc").GetZapLogger()),
	)
}
