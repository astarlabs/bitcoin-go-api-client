package services

import (
	"encoding/json"

	"github.com/astarlabs/bitcoin-go-api-client/dto/result"
	"github.com/astarlabs/bitcoin-go-api-client/http"
)

func APIAddress() result.ServerInfo {
	response, err := httpservice.HTTPRequest(httpservice.GET, "https://astarlabs.github.io/bitcoin-client-server/server-info.json", nil)

	if err != nil {
		return result.ServerInfo{}
	}

	responseResult := result.ServerInfo{}
	err = json.Unmarshal([]byte(response), &responseResult)

	if err != nil {
		return result.ServerInfo{}
	}

	return responseResult
}
