package util

import (
	"Hybrid_Cluster/hcp-apiserver/pkg/util"
	"encoding/json"
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

	//unmarshal Json
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue
}

func PrintOutput(bytes []byte) error {
	var output util.Output
	err := json.Unmarshal(bytes, &output)
	if output.Stderr != nil {
		fmt.Print(string(output.Stderr))
	}
	if output.Stdout != nil {
		fmt.Print(string(output.Stdout))
	}
	return err
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
