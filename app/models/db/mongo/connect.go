package mongo

import (
	"../../../helpers"
	"gopkg.in/mgo.v2"
	"fmt"
)

// MONGO
func Connect() *mgo.Database {
	dbConfig := helpers.GetDatabaseConfig("mongodb")
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
	
	session, err := mgo.Dial(dbString)
	// defer session.Close()
	if err != nil {
	  helpers.ErrorHandler(err)
	}

	db := session.DB(dbname)
	return db
}