package settings

import (
"encoding/json"
"fmt"
"io/ioutil"
"os"
)
type Settings struct {
Api string
}
func GetSettings() Settings{
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//m := new(Dispatch)
	//var m interface{}
	var settings Settings
	json.Unmarshal(file, &settings)
	return settings
}