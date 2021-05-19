package profile

import (
	"context"
	identify "coolcar/proto/identify/gen/go"
	"coolcar/server/rental/profile/dao"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	Mongo  *dao.Mongo
	identify.UnimplementedProfileServiceServer
}

func (s *Service) GetProfile(context.Context, *identify.GetProfileRequest) (*identify.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfile not implemented")
}
func (s *Service) SubmitProfile(context.Context, *identify.Identity) (*identify.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitProfile not implemented")
}
func (s *Service) ClearProfile(context.Context, *identify.ClearProfileRequest) (*identify.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearProfile not implemented")
}
