package db

import (
	"RestApi/interal/user"
	"RestApi/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d db) Create(ctx context.Context, user user.User) (string, error) {
	d.logger.Debug("create user")
	result, err := d.collection.InsertOne(ctx, user)
	if err != nil {
		return "", fmt.Errorf("failed to create user : %s", err)
	}
	d.logger.Debug("failed to convert objectId")
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	d.logger.Trace(user)
	return "", fmt.Errorf("failed to convert objectId %v", oid)
}

func (d db) FindOne(ctx context.Context, id string) (u user.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to obgectId. Hex :%s", id)
	}
	query := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, query)
	if result.Err() != nil {
		// todo 404
		return u, fmt.Errorf("failed to find one user by id: %s", id)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode (user: %s)", id)
	}
	return u, nil
}

func (d db) Update(ctx context.Context, user user.User) error {

}

func (d db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {

	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
