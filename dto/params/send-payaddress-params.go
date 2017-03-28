package params

type SendPayAddressParams struct {
	Account    string
	User       string
	Password   string
	Address    string
	Amount     float64
	Coin       string
	Test       int
	PrivateKey string
}
