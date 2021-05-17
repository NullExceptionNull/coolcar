package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//const IDFiled = "_id"

const (
	IDFieldName       = "_id"
	UpdateAtFieldName = "updateat"
)

var NewObjID = primitive.NewObjectID

var UpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat" `
}
