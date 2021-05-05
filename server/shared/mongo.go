package mgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//const IDFiled = "_id"

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

type ObjID struct {
	ID primitive.ObjectID `bson:"_id"`
}
