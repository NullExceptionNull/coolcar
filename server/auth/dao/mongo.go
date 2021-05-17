package dao

import (
	"context"
	mongo2 "coolcar/server/shared/mongo"
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

	insertedID := mongo2.NewObjID()

	res := m.col.FindOneAndUpdate(c,
		bson.M{openIDField: openID},
		mongo2.SetOnInsert(bson.M{
			mongo2.IDFieldName: insertedID,
			openIDField:        openID}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After),
	)
	if err := res.Err(); err != nil {
		return "", err
	}
	var row mongo2.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", err
	}
	return row.ID.Hex(), nil
}
