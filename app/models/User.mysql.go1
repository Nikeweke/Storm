/*
*  More about methods GORM: http://doc.gorm.io/crud.html#query
*  
*/

/*
*  Same model will be for sqlite, postgres and mssql, difference only in connection: Mysql(), Sqlite(), Postgres(), Mssql()
*/

/*
*  How to use in controllers:
*	 1) import "../models"
*	 2) var UserModel models.User - define model variable
*	 3) call methods 
*
*/
package models

import "time" 

type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	Created_at time.Time
	Updated_at time.Time
}

/**
*  Get all users
*/
func (this User) GetAllUsers() []User {
	db := Mysql()    
	defer db.Close()

	var users []User  // define array of User types
	db.Find(&users)   // mutatiting array by reference 
	
	return users
}


/**
*  Get user by Id
*/
func (this User) GetUserById(id string) User {
	var user User
	Mysql().Where("id = ?", id).First(&user)  

	return user
}


/**
*  Create user
*/
func (this User) CreateUser() {
	db := Mysql()    
	defer db.Close()

	user := User{Name: "Jinzhu", Email: "meta.ua", Created_at: time.Now()}

	// check if record new
  db.NewRecord(user) // => returns `true` as primary key is blank
	db.Create(&user)
	db.NewRecord(user) // => return `false` after `user` created
}


/**
*  Delete user
*/
func (this User) DeleteUser(id string) {
	// Mysql().Where("name LIKE ?", "%John%").Delete(User{})
	Mysql().Where("id = ?", id).Delete(User{})
}


/**
*  Update user
*/
func (this User) UpdateUser(id string) {
	db := Mysql()    
	defer db.Close()
	
	var user User
	db.Where("id = ?", id).First(&user)  

  user.Name = "jinzhu 2"
	user.Email = "gmail.com"
	
  db.Save(&user)
}

