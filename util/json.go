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

// func BytesToJson(bytes []byte, output interface{}) (*gabs.Container, interface{}, error) {
// 	json.Unmarshal(bytes, &output)
// 	// e := reflect.ValueOf(&output).Elem()
// 	jsonParsed, err := gabs.ParseJSON(bytes)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	return jsonParsed, output, err
// }

// func PrintMessage(bytes []byte, niloutput interface{}) error {
// 	output := niloutput
// 	jsonParsed, output, err := BytesToJson(bytes, output)
// 	// fmt.Println(output)
// 	if err != nil {
// 		return err
// 	}
// 	outerrmsg := jsonParsed.Path("Message_").Data()
// 	inerrmsg := jsonParsed.Path("ErrStatus").Path("message").Data()
// 	// e := reflect.ValueOf(&output).Elem()
// 	fmt.Println(reflect.TypeOf(output))
// 	if reflect.ValueOf(output).IsNil() {
// 		if outerrmsg != nil {
// 			fmt.Println(outerrmsg)
// 		} else if inerrmsg != nil {
// 			fmt.Println(inerrmsg)
// 		}
// 	} else {
// 		fmt.Println(output)
// 	}
// 	return nil
// }

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
