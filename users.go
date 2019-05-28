package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var usersCollection = InitDB().Collection("users")

type User struct {
	ID      primitive.ObjectID 	`bson:"_id" json:"id"`
	Name	string 				`bson:"name" json:"name"`
	Secret string             	`bson:"secret" json:"secret"`
}

func (u User) Create() (User, error) {
	u.ID = primitive.NewObjectID()
	u.Secret = generateSecret()
	_, err := usersCollection.InsertOne(context.TODO(), u)
	hexSecret, err := hex.DecodeString(u.Secret)
	if err != nil {
		return u, err
	}
	Secrets[u.ID.String()] = hexSecret
	return u, err
}

func (u User) Exist() bool {
	count, _ := usersCollection.CountDocuments(context.TODO(), bson.M{"_id": u.ID})
	if count != 0 {
		return true
	}
	return false
}

func (u User) Delete() error {
	err := usersCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": u.ID}).Err()
	delete(Secrets, u.ID.String())
	return err
}

func (u User) GetAll() ([]User, error) {
	var users []User
	ctx := context.TODO()
	cur, err := usersCollection.Find(ctx, bson.D{})
	if err != nil { 
		return users, err 
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var res User
		err := cur.Decode(&res)
		if err != nil { 
			return users, err 
		}
		users = append(users, res)
	}
	if err := cur.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func InitSecrets() error {
	users, err := User{}.GetAll()
	if err != nil {
		return err
	}
	for _, u := range users {
		hexSecret, err := hex.DecodeString(u.Secret)
		if err != nil {
			return err
		}
		Secrets[u.ID.String()] = hexSecret
	}
	return err
}


func generateSecret() string {
	const len = 16
	bytes := make([]byte, len)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}