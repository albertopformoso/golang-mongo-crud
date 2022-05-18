package user_repository

import (
	"context"
	"golang-mongo/config"
	m "golang-mongo/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = config.ConnectDB("collection01")
var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)

func Create(user m.User) error {

	var err error

	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	defer cancel()

	return nil
}

func Read() (m.Users, error) {

	var users m.Users

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {
		var user m.User
		err = cur.Decode(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func Update(user m.User, userId string) error {

	var err error

	oid, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func Delete(userId string) error {

	oid, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
