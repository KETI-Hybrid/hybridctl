package util

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"
)

func OpenAndReadJsonFile(jsonFileName string) []byte {
	jsonFile, _ := os.Open(jsonFileName)

	// if err != nil {
	// 	fmt.Println(err)
	// }
	defer jsonFile.Close()

	//unmarshalling Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func PrintErrMsg(bytes []byte) error {
	jsonParsed, err := gabs.ParseJSON(bytes)
	if err != nil {
		return nil
	}
	outerrmsg := jsonParsed.Path("Message_").Data()
	inerrmsg := jsonParsed.Path("ErrStatus").Path("message").Data()
	if outerrmsg != nil {
		fmt.Println(outerrmsg)
	} else if inerrmsg != nil {
		fmt.Println(inerrmsg)
	} else {
		return nil
	}
	return nil
}
