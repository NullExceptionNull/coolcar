package auth

import (
	"context"
	authpb "coolcar/proto/auth/gen/go"
	"coolcar/server/auth/dao"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.uber.org/zap"
)

type Service struct {
	authpb.UnimplementedAuthServiceServer
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	TokenGenerator TokenGenerator
	OPenIDResolver OPenIDResolver
	TokenExpire    time.Duration
}

type OPenIDResolver interface {
	Resolve(code string) (string, error)
}

type TokenGenerator interface {
	GenerateToken(accountID string, expireTime time.Duration) (string, error)
}

func (s *Service) Login(c context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	s.Logger.Info("Received code", zap.String("code", request.Code))
	openID, err := s.OPenIDResolver.Resolve(request.Code)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Get openId error : %v", err)
	}
	account, err := s.Mongo.ResolveAccountID(c, openID)

	if err != nil {
		s.Logger.Error("cannot resolve the account id", zap.Error(err))
		return nil, status.Error(codes.Internal, "account error")
	}

	token, err := s.TokenGenerator.GenerateToken(account, s.TokenExpire)

	if err != nil {
		s.Logger.Error("GenerateToken token error account id", zap.Error(err))
	}
	s.Logger.Info("Get OpenID :  ", zap.String("openID", openID))
	return &authpb.LoginResponse{
		Token:     token,
		ExpiresIn: int32(s.TokenExpire),
	}, nil
}
