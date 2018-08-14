package models

import (
	"../../../helpers"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

// MYSQL
func Connect() *gorm.DB {
	dbConfig := helpers.GetDatabaseConfig("mysql")
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]

	dbString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
	  helpers.ErrorHandler(err)
	}
	return db
}
