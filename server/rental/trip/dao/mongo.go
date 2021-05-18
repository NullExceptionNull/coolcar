package dao

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	mgutil "coolcar/server/shared/mongo"
	"coolcar/server/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

const (
	tripField      = "trip"
	accountIDField = tripField + ".accountid"
	statusField    = tripField + ".status"
)

type TripRecord struct {
	mgutil.IDField        `bson:"inline"`
	mgutil.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.UpdatedAt = mgutil.UpdatedAt()
	r.ID = mgutil.NewObjID()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
func (m *Mongo) GetTrip(c context.Context, id string, accountID string) (*TripRecord, error) {

	objId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	result := m.col.FindOne(c, bson.M{
		mgutil.IDFieldName: objId,
		accountIDField:     accountID,
	})

	if err := result.Err(); err != nil {
		return nil, err
	}
	var tr TripRecord

	err = result.Decode(&tr) //给一个地址 就开始给这个地址写数据

	if err != nil {
		return nil, err
	}
	return &tr, nil
}

func (m *Mongo) GetTrips(c context.Context, accountID string, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountIDField: accountID,
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFICED {
		filter[statusField] = status
	}
	res, err := m.col.Find(c, filter)

	if err != nil {
		return nil, err
	}

	var trips []*TripRecord

	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}

func (m *Mongo) UpdateTrip(c context.Context, tripId string, accountId string, updateAt int64, trip *rentalpb.Trip) error {

	tid, _ := objid.FromID(trip)

	_, err := m.col.UpdateOne(c, bson.M{
		mgutil.IDFieldName:       tid,
		accountIDField:           accountId,
		mgutil.UpdateAtFieldName: updateAt,
	}, mgutil.Set(bson.M{
		tripField:                trip,
		mgutil.UpdateAtFieldName: updateAt,
	}))
	return err
}
