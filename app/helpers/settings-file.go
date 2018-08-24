package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	color "github.com/fatih/color"
)

/*
|--------------------------------------------------------------------------
|  Get settings file and get "databases" section
|--------------------------------------------------------------------------
*/
func GetDatabaseSettings(dbName string) StringArray {
  var settings StringArray = parseSettingsFile("settings.json")

	databases := settings["databases"].(map[string]interface{}) 
	dbConfig  := databases[dbName].(map[string]interface{})

	return dbConfig
}


/*
|--------------------------------------------------------------------------
|  Get settings file and get "databases" section
|--------------------------------------------------------------------------
*/
func GetSettings(name string) interface{} {
	var settings StringArray = parseSettingsFile("settings.json")
	return settings[name] 
}


/*
|--------------------------------------------------------------------------
|  Get settings and parse it
|--------------------------------------------------------------------------
*/
func parseSettingsFile(filename string) StringArray {

	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	color.New(color.FgCyan).Add(color.Bold).Print(strings.Title(filename))
	color.New(color.FgGreen).Add(color.Bold).Print(" successfully loaded \n")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var config StringArray

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &config)

	return config
}




