package hwcloudoperator

import (
	"github.com/gaogj/hw-cloud-operator/resources"
	"github.com/gaogj/hw-cloud-operator/utils"
)

type Operator struct {
	Resources *Api.Resources
}

func NewOperator(cfg *utils.Config) (*Operator) {
	Api.InitHttpClient(cfg)
	return &Operator{
		Resources: Api.NewResources(),
	}
}