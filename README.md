# Bitcoin Go API Client

For more information, please visit [http://www.astarlabs.com](http://www.astarlabs.com)

## Installation
```
    glide install github.com/astarlabs/bitcoin-go-api-client
```

## Getting Started

Please follow the [installation](#installation) instruction and execute the following Java code:

```go

import "github.com/astarlabs/bitcoin-go-api-client/services"
import "github.com/astarlabs/bitcoin-go-api-client/dto/params"
import "fmt"

type SearchSample struct {
}

func (search SearchSample) SearchByRegisteredId() {

	thisSearch := new(services.Search)

	params := new(params.SearchByRegisteredIDParams)
	params.Account = "15"
	params.ID = "333"
	params.User = "test"
	params.Password = "test"
	params.PrivateKey = "***************************************************"

	transaction, err := thisSearch.SearchByRegisteredID(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transaction)

}

```

## Documentation for API Endpoints

All URIs are relative to fulladdress parameter: [server-info.json](https://github.com/astarlabs/bitcoin-client-server/blob/master/server-info.json)

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SearchApi* | [**SearchByAPIID**](docs/SearchApi.md#searchbyapiid) | **Post** /search/registered/id | Get transaction informations by API ID
*SearchApi* | [**SearchByContent**](docs/SearchApi.md#searchbycontent) | **Post** /search/registered/content | Get transaction informations by file or string content
*SearchApi* | [**SearchByHash**](docs/SearchApi.md#searchbyhash) | **Post** /search/registered/hash | Get transaction informations by file or string hash
*SendApi* | [**SendFile**](docs/SendApi.md#sendfile) | **Post** /send/opreturn/base64 | Send file hash to bitcoin blockchain
*SendApi* | [**SendHash**](docs/SendApi.md#sendhash) | **Post** /send/opreturn/hash | Send hash to bitcoin blockchain
*SendApi* | [**SendPayAddress**](docs/SendApi.md#sendpayaddress) | **Post** /send/payaddress | Send a value for address
*SendApi* | [**SendString**](docs/SendApi.md#sendstring) | **Post** /send/opreturn/string | Send string to bitcoin blockchain


## Documentation For Models

 - [SingleResult](docs/SingleResult.md)
 - [Transaction](docs/Transaction.md)


## Author

contato@astarlabs.com

