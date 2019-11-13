package utils

type Config struct {
	AccessKey string
	SecretAccessKey string
	AccountId string
	Endpoints map[string]Endpoint
}

type Endpoint struct {
	Host string
	ProjectId string
}

