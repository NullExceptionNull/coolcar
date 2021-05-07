package server

import (
	"coolcar/server/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServerConfig struct {
	Name              string
	Addr              string
	Logger            *zap.Logger
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
}

func RunGrpcServer(c *GrpcServerConfig) error {
	nameField := zap.String("name", c.Name)

	listen, err := net.Listen("tcp", c.Addr)
	if err != nil {
		c.Logger.Fatal("cannot listen", nameField, zap.Error(err))
	}

	var ops []grpc.ServerOption
	if c.AuthPublicKeyFile != "" {
		interceptor, err := auth.Interceptor(c.AuthPublicKeyFile)
		if err != nil {
			c.Logger.Fatal("cannot authenticate")
		}
		option := grpc.UnaryInterceptor(interceptor)
		ops = append(ops, option)
	}
	server := grpc.NewServer(ops...)

	c.RegisterFunc(server)

	return server.Serve(listen)

}
