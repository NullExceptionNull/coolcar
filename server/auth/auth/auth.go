package auth

import (
	"context"
	authpb "coolcar/proto/auth/gen/go"

	"go.uber.org/zap"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	Logger zap.Logger
}

func (s *Service) Login(c context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("Received code", zap.String("code", request.Code))
	return nil, nil
}
