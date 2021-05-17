package dao

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongo_CreateTrip(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		t.Fatalf("Cannot connect to MongoDB")
	}
	m := NewMongo(client.Database("coolcar"))

	trip, err := m.CreateTrip(ctx, &rentalpb.Trip{
		AccountId: "accountId1",
		CarId:     "car1",
		Start: &rentalpb.LocationStatus{
			PoiName: "StartPoint",
			Location: &rentalpb.Location{
				Latitude:  30,
				Longitude: 120,
			},
		},
		End: &rentalpb.LocationStatus{
			PoiName:  "EndPoint",
			FeeCent:  1000,
			KmDriven: 35,
			Location: &rentalpb.Location{
				Longitude: 115,
				Latitude:  35,
			},
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Errorf("cannot create trip : %v", err)
	}

	record, err := m.GetTrip(ctx, trip.ID.Hex(), "accountId1")

	t.Errorf("Get trip : %+v", record)

	t.Errorf("Res : %+v", trip)
}
