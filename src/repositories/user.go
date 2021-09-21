package repositories

import (
	"barber/src/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var UserCollection IUserCollection

type IUserCollection interface {
	Create(context.Context, *models.User) error
	List(context.Context) ([]*models.User, error)
	FindByID(context.Context, string) (*models.User, error)
	UpdateByID(context.Context, string, *models.User) error
	DeleteByID(context.Context, string) error
}

type userDatabaseHelper struct {
	collection *mongo.Collection
}

func NewUserCollection() IUserCollection {
	return &userDatabaseHelper{
		collection: Database.Collection("users"),
	}
}

func (u *userDatabaseHelper) Create(ctx context.Context, user *models.User) error {
	user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()

	_, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDatabaseHelper) List(ctx context.Context) ([]*models.User, error) {
	cur, err := u.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	users := make([]*models.User, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = cur.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (u *userDatabaseHelper) FindByID(ctx context.Context, id string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user *models.User
	err = u.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userDatabaseHelper) UpdateByID(ctx context.Context, id string, user *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"name":      user.Name,
		"email":     user.Email,
		"provider":  user.Provider,
		"updatedAt": time.Now(),
	}

	var updatedUser *models.User
	err = u.collection.FindOneAndUpdate(ctx, bson.M{"_id": objectID}, bson.M{"$set": update}).Decode(&updatedUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *userDatabaseHelper) DeleteByID(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = u.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	return nil
}
