package server

type ApiKey struct {
	ApiKey string `json:"api_key"`
}

func (k *ApiKey) isValid() bool {
	return medsengerAgentKey == k.ApiKey
}
