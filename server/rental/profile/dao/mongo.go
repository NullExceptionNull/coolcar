package dao

import (
	"context"
	identify "coolcar/proto/identify/gen/go"
	mgutil "coolcar/server/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	accountIDField = "accountid"
	profileField   = "profile"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

type ProfileRecord struct {
	AccountID string            `bson:"accountid"`
	Profile   *identify.Profile `bson:"profile"`
}

func (m *Mongo) GetProfile(c context.Context, accountID string) (*identify.Profile, error) {
	result := m.col.FindOne(c, byAccountID(accountID))
	if err := result.Err(); err != nil {
		return nil, err
	}
	var pr ProfileRecord

	_ = result.Decode(&pr)

	return pr.Profile, nil
}

func (m *Mongo) UpdateProfile(c context.Context, accountID string, p *identify.Profile) error {
	_, err := m.col.UpdateOne(c, byAccountID(accountID), mgutil.Set(bson.M{
		profileField: p,
	}), options.Update().SetUpsert(true))
	return err
}

func byAccountID(accountID string) bson.M {
	return bson.M{
		accountIDField: accountID,
	}
}
