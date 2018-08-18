package models

import (
	"../../../helpers"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
)


// POSTGRES
func Connect() *gorm.DB {
	dbConfig := helpers.GetDatabaseSettings("postgres")
	
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	host     := dbConfig["host"]
	port     := dbConfig["host"]

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
	db, err := gorm.Open("postgres", dbString)
	if err != nil {
	  helpers.ErrorHandler(err)
	}
  return db 
}

