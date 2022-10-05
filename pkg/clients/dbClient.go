package clients

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type DBClient interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}) (*mongo.SingleResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error)
	InsertMany(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error)
}

type dbClient struct {
	dbclient *mongo.Collection
}

func (d *dbClient) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	insertRes, err := d.dbclient.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return insertRes, nil
}

func (d *dbClient) InsertMany(ctx context.Context, documents []interface{}) (*mongo.InsertManyResult, error) {
	insertRes, err := d.dbclient.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	return insertRes, nil
}

func (d *dbClient) FindOne(ctx context.Context, filter interface{}) (*mongo.SingleResult, error) {
	findRes := d.dbclient.FindOne(ctx, filter)
	if findRes.Err() != nil && findRes.Err() != mongo.ErrNoDocuments {
		return nil, findRes.Err()
	}
	return findRes, nil
}

func (d *dbClient) UpdateOne(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	updateRes, err := d.dbclient.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return updateRes, nil
}
