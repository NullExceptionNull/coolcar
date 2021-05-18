package trip

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip/dao"
	"coolcar/server/shared/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger *zap.Logger
	Mongo  *dao.Mongo
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
func (s *Service) UpdateTrip(c context.Context, req *rentalpb.UpdateTripRequest) (*rentalpb.Trip, error) {

	aid, err := auth.CidFromContext(c)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	tr, err := s.Mongo.GetTrip(c, req.Id, aid)

	if req.Current != nil {
		tr.Trip.Current = s.calcCurrentStatus(tr.Trip, req.GetCurrent())
	}
	//如果行程已结束 更新
	if req.EndTrip {
		tr.Trip.End = tr.Trip.Current
		tr.Trip.Status = rentalpb.TripStatus_FINISHED
	}

	s.Mongo.UpdateTrip(c, req.Id, aid, tr.UpdatedAt, tr.Trip)

	return nil, nil
}

//计算当前的行程状态
func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
