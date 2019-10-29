package hwcloudoperator

import (
	"fmt"
	"github.com/gaogj/hw-cloud-operator/utils"
	"io/ioutil"
	"testing"
)

func TestVPCNewOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
		AccessKey: "SI85FNG***LMV3RUH",
		SecretAccessKey: "0iskcyceY0****TEv8p9Iv9B1H3kTRd",
		ProjectId: "4e7b81******02a00458",
	})

	res,err := operator.Resources.VPC.Get(operator.Endpoints["bj1"],
		operator.Resources.VPC.Get.WithResourceId("11016528-83***-89620a674a4c"),
	)

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestNewSubnetOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
		AccessKey: "SI85FNG***LMV3RUH",
		SecretAccessKey: "0iskcyceY0****TEv8p9Iv9B1H3kTRd",
		ProjectId: "4e7b81******02a00458",
	})

	res,err := operator.Resources.VPC.GetSubnet(operator.Endpoints["bj1"],operator.Resources.VPC.GetSubnet.WithVPCId("11016528-834a****-89620a674a4c"))

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}