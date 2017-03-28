package services

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/astarlabs/go-bitcoin-api-client/dto/params"
	"github.com/astarlabs/go-bitcoin-api-client/dto/result"
	"github.com/astarlabs/go-bitcoin-api-client/http"
	"github.com/regcostajr/pktoken"
)

type Search struct {
}

// SearchByRegisteredID ...
func (search Search) SearchByRegisteredID(param params.SearchByRegisteredIDParams) (*result.Transaction, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("password", param.Password)
	form.Add("id", param.ID)

	response, err := httpservice.HTTPRequest(httpservice.POST, APIAddress+"/search/registered/id", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.Transaction{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}

// SearchByRegisteredContent ...
func (search Search) SearchByRegisteredContent(param params.SearchByRegisteredContentParams) (*result.Transaction, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("password", param.Password)
	form.Add("content", param.Content)

	response, err := httpservice.HTTPRequest(httpservice.POST, APIAddress+"/search/registered/content", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.Transaction{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}

// SearchByRegisteredHash ...
func (search Search) SearchByRegisteredHash(param params.SearchByRegisteredHashParams) (*result.Transaction, error) {

	form := url.Values{}

	token, err := pktoken.SignMessage(param.PrivateKey)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	form.Add("token", token)
	form.Add("account", param.Account)
	form.Add("user", param.User)
	form.Add("password", param.Password)
	form.Add("hash", param.Hash)

	response, err := httpservice.HTTPRequest(httpservice.POST, APIAddress+"/search/registered/hash", form)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := result.Transaction{}

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}
