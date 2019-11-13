package hwcloudoperator

import (
	"fmt"
	"github.com/gaogj/hw-cloud-operator/utils"
	"io/ioutil"
	"testing"
)
//
//func TestVPCNewOperator(t *testing.T) {
//	operator := NewOperator(&utils.Config{
//		AccessKey: "*****",
//		SecretAccessKey: "*****",
//		ProjectId: "*****",
//	})
//
//	res,err := operator.Resources.VPC.Get(operator.Endpoints["bj1"])
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	body,err := ioutil.ReadAll(res.Body)
//	fmt.Println(string(body))
//}
////
//func TestNewSubnetOperator(t *testing.T) {
//	operator := NewOperator(&utils.Config{
//		AccessKey: "*****",
//		SecretAccessKey: "*****",
//		ProjectId: "*****",
//	})
//
//	res,err := operator.Resources.VPC.GetSubnet(operator.Endpoints["bj1"])
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	body,err := ioutil.ReadAll(res.Body)
//	fmt.Println(string(body))
//}

func TestGetUserOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
		AccessKey: "*****",
		SecretAccessKey: "*****",
		ProjectId: "*****",
	})

	res,err := operator.Resources.IAM.GetUser(operator.Endpoints["bj1"])

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Header["X-Request-Id"])
	fmt.Println(string(body))
}
