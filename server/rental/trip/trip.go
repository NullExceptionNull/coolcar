package trip

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"go.uber.org/zap"
)

type Service struct {
	Logger *zap.Logger
	rentalpb.UnimplementedTripServiceServer
}

func (s *Service) CreateTrip(context.Context, *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	return nil, nil
}
func (s *Service) GetTrip(context.Context, *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {
	return nil, nil

}
func (s *Service) GetTrips(context.Context, *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {
	return nil, nil

}
func (s *Service) UpdateTrip(context.Context, *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {
	return nil, nil
}
