package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astarlabs/bitcoin-go-api-client/dto/result"
)

func APIAddress() string {
	response, err := http.Get("https://astarlabs.github.io/bitcoin-client-server/server-info.json")
	response.Header.Add("Content-Type", "application/json")

	defer response.Body.Close()

	if err != nil {
		fmt.Println(err)
		return ""
	}

	responseResult := []result.ServerInfo{}

	body, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &responseResult)

	return responseResult[0].FullAddress
}
