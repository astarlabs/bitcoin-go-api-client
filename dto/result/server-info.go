package result

type ServerInfo struct {
	FullAddress string `json:"fulladdress"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Version     string `json:"v2"`
	Protocol    string `json:"https"`
}
