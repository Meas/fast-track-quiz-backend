package helpers

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJSON(fileName string, questions interface{}) {
	jsonFile, _ := os.Open(fileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &questions)
	defer jsonFile.Close()
}
