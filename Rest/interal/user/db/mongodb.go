package db

import (
	user2 "RestApi/Rest/interal/user"
	"RestApi/Rest/pkg/logging"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user2.User) (string, error) {
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

func (d *db) FindOne(ctx context.Context, id string) (u user2.User, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, fmt.Errorf("failed to convert hex to obgectId. Hex :%s", id)
	}
	query := bson.M{"_id": oid}
	result := d.collection.FindOne(ctx, query)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			// TODO errEntityNotFound
		}
		return u, fmt.Errorf("failed to find one user by id: %s", id)
	}
	if err = result.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode (user: %s)", id)
	}
	return u, nil
}

func (d *db) FindAll(ctx context.Context) (u []user2.User, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		return u, fmt.Errorf("failed to all users %v", err)
	}

	if err = cursor.All(ctx, &u); err != nil {
		return u, fmt.Errorf("failde to read all documents from cursor")
	}
	if err = cursor.Decode(&u); err != nil {
		return u, fmt.Errorf("failed to decode all users")
	}
	return u, nil
}

func (d *db) Update(ctx context.Context, user user2.User) error {
	objectId, err := primitive.ObjectIDFromHex(user.Id)
	if err != nil {
		return fmt.Errorf("failed to convert user Id to ObjectId, id = %s")
	}
	filter := bson.M{"_id": objectId}
	userBytes, err := bson.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marhsal user. error: %v", err)
	}
	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		return fmt.Errorf("failed to unmarshal. err:%v", err)
	}
	delete(updateUserObj, "_id")
	update := bson.M{
		"$set": updateUserObj,
	}
	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failded to execute update user query. err: %v", err)
	}
	if result.MatchedCount == 0 {
		// TODO errEntityNotFound
		return fmt.Errorf("not found if  for update")
	}
	d.logger.Tracef("Matched %d document and Modified %d dock", result.MatchedCount, result.ModifiedCount)
	return nil
}

func (d *db) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("failde to cinvert user ID to ObjectId. ID=%s", id)
	}
	filter := bson.M{"_id": objectId}
	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		// TODO errEntyti
		return fmt.Errorf("not found id for delite")
	}

	return nil
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user2.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
