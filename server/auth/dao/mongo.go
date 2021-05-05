package dao

import (
	"context"
	mgo "coolcar/server/shared"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIDField = "open_id"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountID(c context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(c,
		bson.M{openIDField: openID},
		mgo.Set(bson.M{openIDField: openID}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if err := res.Err(); err != nil {
		return "", err
	}
	var row mgo.ObjID
	err := res.Decode(&row)
	if err != nil {
		return "", err
	}
	return row.ID.Hex(), nil
}
