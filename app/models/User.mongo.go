/*
*  More about Mongo library for golang: https://labix.org/mgo
*  
*/

/*
*  How to use in controllers:
*	 1) import "../models"
*	 2) define model variable -  var UserModel models.User 
*	 3) call methods          -  users := UserModel.GetAllUsers()
*/
package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"./db/mongo"
)

type User struct {
	Email      string
	Password   string
  Id         bson.ObjectId `json:"id" bson:"_id,omitempty"`
}

/**
*  Get all users
*/
func (this User) GetAllUsers() []User {
  mongoSession, DB_NAME := mongo.Connect()
	defer mongoSession.Close()
	
	var users []User  // define array of User types
	userCollection := mongoSession.DB(DB_NAME).C("users")

	// get one item
	err := userCollection.Find(bson.M{}).All(&users)
  if err != nil {
		fmt.Println(err)
	}

	return users
}


/**
*  Get user by Id
*/
func (this User) GetUserById(id string) User {
	mongoSession, DB_NAME := mongo.Connect()
	defer mongoSession.Close()   
  
	var user User 
	userCollection := mongoSession.DB(DB_NAME).C("users")

	// get one item
	err := userCollection.Find(bson.M{"_id": id}).One(&user)
  if err != nil {
		fmt.Println(err)
	}
	
	return user
}


/**
*  Create user
*/
func (this User) CreateUser() {
	mongoSession, DB_NAME := mongo.Connect()
	defer mongoSession.Close()   

	userCollection := mongoSession.DB(DB_NAME).C("users")
	user           := User{Email: "gmail.com", Password: "2233221"} 

	err := userCollection.Insert(user)
  if err != nil {
		fmt.Println(err)
	}
}


/**
*  Delete user
*/
func (this User) DeleteUser(id string) {
	mongoSession, DB_NAME := mongo.Connect()
	defer mongoSession.Close()  

	userCollection := mongoSession.DB(DB_NAME).C("users")

	err := userCollection.Remove(bson.M{"_id": id})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
	}
}


/**
*  Update user
*/
func (this User) UpdateUser(id string) {
	mongoSession, DB_NAME := mongo.Connect()
	defer mongoSession.Close()  

	userCollection := mongoSession.DB(DB_NAME).C("users") 
	
	err := userCollection.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"email": "new Name"}})
	if err != nil {
		fmt.Printf("update fail %v\n", err)
	}
}





