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
func loadFile(filename string) []byte{
	file, e := ioutil.ReadFile(filename)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return file
}
func GetSettings() Settings{
	var settings Settings
	json.Unmarshal(loadFile("./config.json"), &settings)
	return settings
}