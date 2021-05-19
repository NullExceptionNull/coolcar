package trip

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"coolcar/server/rental/trip/dao"
	"coolcar/server/shared/auth"
	"coolcar/server/shared/id"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	Logger         *zap.Logger
	Mongo          *dao.Mongo
	ProfileManager ProfileManager
	CarManager     CarManager
	POIManager     POIManager
	rentalpb.UnimplementedTripServiceServer
}

type CarManager interface {
	Verify(context.Context, id.CarID) error
	UnLock(context.Context, id.CarID) error
}

// ProfileManager defines the ACL
type ProfileManager interface {
	Verify(ctx context.Context, accountID string) (id.IdentityId, error)
}

type POIManager interface {
	Resolve(context.Context, *rentalpb.Location) (string, error)
}

func (s *Service) CreateTrip(c context.Context, req *rentalpb.CreateTripRequest) (*rentalpb.TripEntity, error) {
	accountID, err := auth.CidFromContext(c)

	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, err.Error())
	}
	//验证驾驶者身份
	identifyId, err := s.ProfileManager.Verify(c, accountID)

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}
	//检查车辆状态是不是可以出租
	carId := id.CarID(req.CarId)

	err = s.CarManager.Verify(c, carId)

	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, err.Error())
	}

	startName, _ := s.POIManager.Resolve(c, req.Start)

	ls := &rentalpb.LocationStatus{
		Location: req.Start,
		PoiName:  startName,
	}

	//创建行程 开始计费
	tr, err := s.Mongo.CreateTrip(c, &rentalpb.Trip{
		AccountId:  accountID,
		CarId:      carId.String(),
		IdentityId: identifyId.String(),
		Status:     rentalpb.TripStatus_IN_PROGRESS,
		Start:      ls,
		Current:    ls,
	})
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, "")
	}
	//车辆开锁
	go func() {
		err = s.CarManager.UnLock(c, carId)
		if err != nil {
			s.Logger.Error("cannot unlock car", zap.Error(err))
		}
	}()

	return &rentalpb.TripEntity{
		Id:   tr.ID.Hex(),
		Trip: tr.Trip,
	}, nil
}
func (s *Service) GetTrip(c context.Context, req *rentalpb.GetTripRequest) (*rentalpb.Trip, error) {

	aid, err := auth.CidFromContext(c)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	trip, _ := s.Mongo.GetTrip(c, req.Id, aid)

	return trip.Trip, nil

}
func (s *Service) GetTrips(c context.Context, req *rentalpb.GetTripsRequest) (*rentalpb.GetTripsResponse, error) {

	aid, err := auth.CidFromContext(c)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}

	trips, _ := s.Mongo.GetTrips(c, aid, req.Status)

	var res = new(rentalpb.GetTripsResponse)

	for _, tr := range trips {
		res.Trips = append(res.Trips, &rentalpb.TripEntity{
			Id:   tr.ID.String(),
			Trip: tr.Trip,
		})
	}
	return res, nil
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

	return tr.Trip, nil
}

//计算当前的行程状态
func (s *Service) calcCurrentStatus(trip *rentalpb.Trip, cur *rentalpb.Location) *rentalpb.LocationStatus {
	return nil
}
