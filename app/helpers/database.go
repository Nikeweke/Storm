package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	color "github.com/fatih/color"
)

/*
|--------------------------------------------------------------------------
|  Get settings.json and parse to map 
|--------------------------------------------------------------------------
*/
func GetDatabaseConfig(dbName string) map[string]interface{} {

	// Open our jsonFile
	jsonFile, err := os.Open("settings.json")
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

