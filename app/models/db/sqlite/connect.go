package models

import (
	"../../../helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)


// SQLITE
func Connect() *gorm.DB {
	dbConfig := helpers.GetDatabaseConfig("sqlite")
	dbString := dbConfig["db_name"]
	
	db, err := gorm.Open("sqlite3", dbString)
	if err != nil {
	  helpers.ErrorHandler(err)
	}
  return db 
}

