package utils

type Config struct {
	AccessKey string
	SecretAccessKey string
	AccountId string
}

var Endpoints map[string]interface{} = map[string]interface{}{
	"bj1": map[string]string{
		"host": "cn-north-1.myhuaweicloud.com",
		"projectId": "4e7b816ca6524ef3a0700dba02a00458",
	},
}