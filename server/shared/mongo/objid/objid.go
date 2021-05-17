package objid

import (
	"coolcar/server/shared/id"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

func ToTripID(oid primitive.ObjectID) id.TripID {
	return id.TripID(oid.Hex())
}
