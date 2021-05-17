package dao

import (
	"context"
	rentalpb "coolcar/proto/rental/gen/go"
	mongo2 "coolcar/server/shared/mongo"
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
)

type TripRecord struct {
	mongo2.IDField        `bson:"inline"`
	mongo2.UpdatedAtField `bson:"inline"`
	Trip                  *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(c context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.UpdatedAt = mongo2.UpdatedAt()
	r.ID = mongo2.NewObjID()
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
		mongo2.IDFieldName: objId,
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
