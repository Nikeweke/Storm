package models

import (
	color "github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"
	"gopkg.in/mgo.v2"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// MONGO
func Mongo() *mgo.Session {
	dbConfig := getDatabaseConfig("mongodb")
	dbString := ""

	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	port     := dbConfig["port"]
	host     := dbConfig["host"]

	if user == "" || password == ""  {
		dbString = fmt.Sprintf("mongodb://%s/%s:%s", host, dbname, port)
	} else {
    dbString = fmt.Sprintf("mongodb://%s:%s@%s/%s:%s", user, password, host, dbname, port)
	}
	
	db, err := mgo.Dial(dbString)
	if err != nil {
    errorHandler(err)
	}

	return db
}

// MYSQL
func Mysql() *gorm.DB {
	dbConfig := getDatabaseConfig("mysql")
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]

	dbString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)
	db, err := gorm.Open("mysql", dbString)
	if err != nil {
	  errorHandler(err)
	}
	return db
}

// SQLITE
func Sqlite() *gorm.DB {
	dbConfig := getDatabaseConfig("sqlite")
	dbString := dbConfig["db_address"]
	
	db, err := gorm.Open("sqlite3", dbString)
	if err != nil {
	  errorHandler(err)
	}
  return db 
}

// POSTGRES
func Postgres() *gorm.DB {
	dbConfig := getDatabaseConfig("postgres")
	
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	host     := dbConfig["host"]
	port     := dbConfig["host"]

	dbString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", dbString)
	if err != nil {
	  errorHandler(err)
	}
  return db 
}

// MSSQL
func Mssql() *gorm.DB {
	dbConfig := getDatabaseConfig("mssql")
	
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	host     := dbConfig["host"]
	port     := dbConfig["host"]

	dbString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, password, host, port, dbname)
  db, err := gorm.Open("mssql", dbString)
	if err != nil {
	  errorHandler(err)
	}
  return db 
}

/**
*  Get config.json and parse to map 
*/
func getDatabaseConfig(dbName string) map[string]interface{} {

	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Print("Config.json ")
	color.New(color.FgGreen).Add(color.Bold).Print("successfully loaded \n")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var config map[string]interface{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &config)

	databases := config["databases"].(map[string]interface{}) 
	dbConfig  := databases[dbName].(map[string]interface{})

	return dbConfig
}

/**
*  Error handler
*/
func errorHandler(err error) {
	color.Red("ERROR ====================> ")
	fmt.Println(err)
	color.Red("ERROR ====================> ")
}