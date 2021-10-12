package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func UnmarshalJsonFile(jsonFileName string, target interface{}) {
	jsonFile, err := os.Open(jsonFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	//unmarshalling Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &target)
	fmt.Println(target)
}
