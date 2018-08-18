package models

import (
	"../../../helpers"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mssql"
	"fmt"
)


// MSSQL
func Connect() *gorm.DB {
	dbConfig := helpers.GetDatabaseSettings("mssql")
	
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	host     := dbConfig["host"]
	port     := dbConfig["host"]

	dbString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, dbname)
  db, err := gorm.Open("mssql", dbString)
	if err != nil {
	  helpers.ErrorHandler(err)
	}
  return db 
}

