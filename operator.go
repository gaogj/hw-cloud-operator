package hwcloudoperator

import (
	"github.com/gaogj/hw-cloud-operator/resources"
	"github.com/gaogj/hw-cloud-operator/utils"
)

var Endpoints map[string]string = map[string]string{
	"bj1": "cn-north-1.myhuaweicloud.com",
}

type Operator struct {
	Endpoints map[string]string
	Resources *Api.Resources
}

func NewOperator(cfg *utils.Config) (*Operator) {
	Api.InitHttpClient(cfg)
	return &Operator{
		Endpoints: Endpoints,
		Resources: Api.NewResources(cfg.ProjectId),
	}
}