package main

import (
	"context"
	"errors"
	"github.com/Take-A-Seat/storage"
	"github.com/Take-A-Seat/storage/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func getUserByParam(key string, param string) (models.User, error) {
	var user models.User
	var objID primitive.ObjectID
	var err error

	client, err := storage.ConnectToDatabase(mongoUser, mongoPass, mongoHost,mongoDatabase)
	if err != nil {
		return models.User{},err
	}

	defer storage.DisconnectFromDatabase(client)

	usersCollection := client.Database(mongoDatabase).Collection("Users")

	if key == "_id" {
		objID, err = primitive.ObjectIDFromHex(param)
		err = usersCollection.FindOne(context.TODO(), bson.D{{key, objID}}).Decode(&user)
	} else {
		err = usersCollection.FindOne(context.TODO(), bson.D{{key, param}}).Decode(&user)
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func addUser(user models.User) error {
	client, err := storage.ConnectToDatabase(mongoUser, mongoPass, mongoHost,mongoDatabase)
	defer storage.DisconnectFromDatabase(client)
	if err != nil {
		return err
	}

	usersCollection := client.Database(mongoDatabase).Collection("users")
	//Hashing the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	checkUser, _ := getUserByParam("email", user.Email)


	if checkUser.Id.IsZero(){
		_, err := usersCollection.InsertOne(context.Background(), bson.M{"firstName": user.FirstName,
			"lastName": user.LastName, "email": user.Email, "role": user.Role,"password":hashedPassword})
		if err!=nil{
			return err
		}
	}else{
		return errors.New("Duplicate email")
	}

	return nil

}
