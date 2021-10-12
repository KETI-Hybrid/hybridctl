package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetJson(httpPostUrl string, input interface{}, target interface{}) error {

	jsonData, _ := json.Marshal(&input)
	buff := bytes.NewBuffer(jsonData)
	request, _ := http.NewRequest("POST", httpPostUrl, buff)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()
	jsonDataFromHttp, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Headers:", response.Header)

	return json.Unmarshal(jsonDataFromHttp, target)
	// return response, err
}

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
