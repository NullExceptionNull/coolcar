package trip

import (
	"context"
	"coolcar/proto/rental/gen/go"
	"coolcar/server/shared/auth"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(ctx context.Context, request *rentalpb.CreateTripRequest) (*rentalpb.CreateTripResponse, error) {

	aid, err := auth.CidFromContext(ctx)
	if err != nil {
		return nil, err
	}

	s.Logger.Info("Create Trip ", zap.String("Start", aid))

	return &rentalpb.CreateTripResponse{}, nil
}
