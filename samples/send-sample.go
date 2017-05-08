package samples

import (
	"fmt"

	"github.com/astarlabs/bitcoin-go-api-client/dto/params"
	"github.com/astarlabs/bitcoin-go-api-client/services"
)

type SendSample struct {
}

func (send SendSample) SendOPReturnBase64() {

	thisSend := new(services.Send)

	params := new(params.SendOPReturnBase64Params)
	params.Account = "1"
	params.User = "test"
	params.Password = "test"
	params.Base64 = "dGVzdA=="
	params.Coin = "bitcoin" // bitcoin / litecoin
	params.Test = 1         // 1 test / 0 production
	params.PrivateKey = "***************************************************"

	result, err := thisSend.SendOPReturnBase64(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func (send SendSample) SendOPReturnString() {

	thisSend := new(services.Send)

	params := new(params.SendOPReturnStringParams)
	params.Account = "1"
	params.User = "test"
	params.Password = "test"
	params.String = "test"
	params.Coin = "bitcoin" // bitcoin / litecoin
	params.Test = 1         // 1 test / 0 production
	params.PrivateKey = "***************************************************"

	result, err := thisSend.SendOPReturnString(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}

func (send SendSample) SendPayAddress() {

	thisSend := new(services.Send)

	params := new(params.SendPayAddressParams)
	params.Account = "1"
	params.User = "test"
	params.Password = "test"
	params.Address = "1F1tAaz5x1HUXrCNLbtMDqcw6o5GNn4xqX"
	params.Amount = 0.00010000
	params.Coin = "bitcoin" // bitcoin / litecoin
	params.Test = 1         // 1 test / 0 production
	params.PrivateKey = "***************************************************"

	result, err := thisSend.SendPayAddress(*params)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

}
