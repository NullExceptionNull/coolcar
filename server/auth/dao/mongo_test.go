package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongo_ResolveAccountID(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		t.Fatalf("Cannot connect to MongoDB")
	}
	m := NewMongo(client.Database("coolcar"))
	id, err := m.ResolveAccountID(ctx, "123")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Logf("Get account id %v", id)
	}

}
