package trip

import (
	"context"

	trippb "coolcar/proto/gen/go"
)

// type TripServiceServer interface {
// 	GetTrip(context.Context, *GetTripRequest) (*GetTripResp, error)
// }

// Service is a trip service implementation
type Service struct{}

func (r *Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (resp *trippb.GetTripResp, e error) {

	return &trippb.GetTripResp{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "ABC",
			End:         "EDF",
			DurationSec: 1000,
			FeeCent:     10,
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
