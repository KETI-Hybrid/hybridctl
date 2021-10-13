package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getJson(httpPostUrl string, input interface{}, target interface{}) error {

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
}
