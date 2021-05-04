package trip

import (
	"context"
	trippb "coolcar/proto/trip/gen/go"
)

// Service is a trip service implementation
type Service struct {
	trippb.UnimplementedTripServiceServer
}

func (s Service) GetTrip(ctx context.Context, request *trippb.GetTripRequest) (*trippb.GetTripResp, error) {
	return &trippb.GetTripResp{
		Id: request.Id,
		Trip: &trippb.Trip{
			Start:       "ABC",
			End:         "EDF",
			DurationSec: 1000,
			FeeCent:     10,
			Status:      trippb.TripStatus_IN_PROGRESS,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  31,
					Longitude: 119,
				}, {
					Latitude:  32,
					Longitude: 118,
				},
			},
		},
	}, nil
}
