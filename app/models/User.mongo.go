package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
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
	db := Mongo()    

	var users []User  // define array of User types
	userCollection := db.C("users")

	// get one item
	err := userCollection.Find(bson.M{}).All(&users)
  if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
	
	return users
}


/**
*  Get user by Id
*/
func (this User) GetUserById(id string) User {
	db := Mongo()    
  
	var user User 
	userCollection := db.C("users")

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
	db := Mongo()   

	userCollection := db.C("users")
	user           := User{Email: "gmail.com", Password: "2233221"} 

	err := userCollection.Insert(user)
  if err != nil {
		fmt.Println(err)
	}
}


// /**
// *  Delete user
// */
// func (this User) DeleteUser(id string) {
// 	// Mysql().Where("name LIKE ?", "%John%").Delete(User{})
// 	Mysql().Where("id = ?", id).Delete(User{})
// }


// /**
// *  Update user
// */
// func (this User) UpdateUser(id string) {
// 	db := Mysql()    
// 	defer db.Close()
	
// 	var user User
// 	db.Where("id = ?", id).First(&user)  

//   user.Name = "jinzhu 2"
// 	user.Email = "gmail.com"
	
//   db.Save(&user)
// }





