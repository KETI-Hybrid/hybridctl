package util

import (
	"fmt"
	"io/ioutil"
	"os"
)

func OpenAndReadJsonFile(jsonFileName string) []byte {
	jsonFile, err := os.Open(jsonFileName)

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	//unmarshalling Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}
