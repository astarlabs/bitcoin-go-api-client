package services

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/astarlabs/bitcoin-go-api-client/dto/params"
	"github.com/astarlabs/bitcoin-go-api-client/dto/result"
	"github.com/astarlabs/bitcoin-go-api-client/http"

	"strconv"

	"github.com/regcostajr/pktoken"
)

// Send ...
type Send struct{}

// SendOPReturnBase64 ...
func (send Send) SendOPReturnBase64(param params.SendOPReturnBase64Params) (*result.ResultDTO, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("pass", param.Password)
	form.Add("base64", param.Base64)
	form.Add("coin", param.Coin)
	form.Add("test", strconv.Itoa(param.Test))

	response, err := httpservice.HTTPRequest(APIAddress()+"/send/opreturn/base64", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.ResultDTO{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}

// SendOPReturnString ...
func (send Send) SendOPReturnString(param params.SendOPReturnStringParams) (*result.ResultDTO, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("pass", param.Password)
	form.Add("string", param.String)
	form.Add("coin", param.Coin)
	form.Add("test", strconv.Itoa(param.Test))

	response, err := httpservice.HTTPRequest(APIAddress()+"/send/opreturn/string", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.ResultDTO{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}

// SendPayAddress ...
func (send Send) SendPayAddress(param params.SendPayAddressParams) (*result.ResultDTO, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("pass", param.Password)
	form.Add("address", param.Address)

	convAmmount := strconv.FormatFloat(param.Amount, 'f', 9, 64)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("amount", convAmmount)
	form.Add("coin", param.Coin)
	form.Add("test", strconv.Itoa(param.Test))

	response, err := httpservice.HTTPRequest(APIAddress()+"/send/payaddress", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.ResultDTO{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}
