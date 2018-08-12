package models

import (
	color "github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//  _ "github.com/jinzhu/gorm/dialects/postgres"
	//  _ "github.com/jinzhu/gorm/dialects/mssql"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)


func Mysql() {
	dbConfig := getDatabaseConfig("mysql")
	user     := dbConfig["user"]
	password := dbConfig["password"]
	dbname   := dbConfig["db_name"]
	enabled  := dbConfig["enabled"].(bool)

	if enabled {
		dbString := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", user, password, dbname)

		db, err := gorm.Open("mysql", dbString)
		if err != nil {
			color.Red("ERROR ====================> ")
			fmt.Println(err)
			color.Red("ERROR ====================> ")
		}
		defer db.Close()
	}
}



// func DbInit() *sql.DB{
// 	// Открываем соеденение и чекаем на ошибки
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test")  // yourusername:yourpassword@/yourdatabase"
// 	checkErr(err)

// 	// проверяем соедение с БД
// 	err = db.Ping()
// 	checkErr(err)

// 	return db
// }


// func Mongo() {
// }

// func Sqlite() {
//   db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
//   defer db.Close()
// }

// func Postgres() {
// 	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
// 	defer db.Close()
// }

// func Mssql() {
//   db, err = gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
//   defer db.Close()
// }



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