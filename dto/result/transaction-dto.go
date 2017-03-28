package result

import (
	"time"
)

type Transaction struct {
	CreationDate           time.Time `json:"creationdate"`
	ConfirmationDate       time.Time `json:"confirmationdate"`
	BlockchainCreationDate time.Time `json:"blockchaincreationdate"`
	TransactionHash        string    `json:"txid"`
	Amount                 float64   `json:"amount"`
	Fee                    float64   `json:"fee"`
	TransactionSize        int64     `json:"txsize"`
	TransactionHex         string    `json:"txhex"`
	Data                   string    `json:"data"`
	FeePriority            int64     `json:"feepriority"`
	Confirmations          int64     `json:"confirmations"`
	ErrorMessage           string    `json:"errormessage"`
	Coin                   string    `json:"coin"`
	Test                   int64     `json:"test"`
}
