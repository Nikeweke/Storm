package mongo

import (
	"../../../helpers"
	"gopkg.in/mgo.v2"
	"fmt"
)

var DB_NAME = ""

// MONGO
func Connect() (*mgo.Session, string) {
	dbConfig := helpers.GetDatabaseSettings("mongodb")
	dbString := ""

	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"].(string)
	port     := dbConfig["port"]
	host     := dbConfig["host"]

	if user == "" || password == ""  {
		dbString = fmt.Sprintf("mongodb://%s/%s:%s", host, dbname, port)
	} else {
    dbString = fmt.Sprintf("mongodb://%s:%s@%s/%s:%s", user, password, host, dbname, port)
	}
	
	sessionDB, err := mgo.Dial(dbString)
	if err != nil {
		helpers.ErrorHandler(err)
	}
	
  DB_NAME = dbname 

	return sessionDB, DB_NAME
}