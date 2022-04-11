package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetResponseBody(method string, httpPostUrl string, input interface{}) ([]byte, error) {

	jsonData, _ := json.Marshal(&input)
	buff := bytes.NewBuffer(jsonData)
	request, _ := http.NewRequest(method, httpPostUrl, buff)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err, "GetResponseBody Func")
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	return bytes, err
}

func HTTPPostRequest(input interface{}, httpPostUrl string) []byte {
	bytes, err := GetResponseBody("POST", httpPostUrl, input)
	CheckERR(err)
	return bytes
}

func CheckERR(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
