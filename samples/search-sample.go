package samples

import "github.com/astarlabs/go-bitcoin-api-client/services"
import "github.com/astarlabs/go-bitcoin-api-client/dto/params"
import "fmt"

type SearchSample struct {
}

func (search SearchSample) SearchByRegisteredId() {

	thisSearch := new(services.Search)

	params := new(params.SearchByRegisteredIDParams)
	params.Account = "1"
	params.ID = "216"
	params.User = "test"
	params.Password = "test"
	params.PrivateKey = "***************************************************"

	transaction, err := thisSearch.SearchByRegisteredID(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transaction)

}

func (search SearchSample) SearchByRegisteredContent() {

	thisSearch := new(services.Search)

	params := new(params.SearchByRegisteredContentParams)
	params.Account = "1"
	params.Content = "*****************************************************"
	params.User = "test"
	params.Password = "test"
	params.PrivateKey = "***************************************************"

	transaction, err := thisSearch.SearchByRegisteredContent(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transaction)

}

func (search SearchSample) SearchByRegisteredHash() {

	thisSearch := new(services.Search)

	params := new(params.SearchByRegisteredHashParams)
	params.Account = "1"
	params.Hash = "a310189de0b302b4ee765889643891561828b8252824b99a2ecd8d43670eaec4"
	params.User = "test"
	params.Password = "test"
	params.PrivateKey = "***************************************************"

	transaction, err := thisSearch.SearchByRegisteredHash(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(transaction)

}
