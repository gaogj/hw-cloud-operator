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
//	res,err := operator.Resources.VPC.Get(operator.Endpoints["bj1"],
//		operator.Resources.VPC.Get.WithResourceId("11016528-834a-4ee6-b130-89620a674a4c"))
//
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	body,err := ioutil.ReadAll(res.Body)
//	fmt.Println(string(body))
//}
////
func TestNewSubnetOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
		AccessKey: "RLJDCAVZ0PT80SL3O4HP",
		SecretAccessKey: "bBZ3CdmRWHWs79joudeGJz0ntYmIEaHWzd6a86PN",
		AccountId: "a2f58a7bbf264053ab03375e7dbf9501",
	})

	res,err := operator.Resources.VPC.GetSubnet("bj1")

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

func TestGetUserOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
		AccessKey: "RLJDCAVZ0PT80SL3O4HP",
		SecretAccessKey: "bBZ3CdmRWHWs79joudeGJz0ntYmIEaHWzd6a86PN",
		AccountId: "a2f58a7bbf264053ab03375e7dbf9501",
	})

	res,err := operator.Resources.IAM.GetUser("bj1",
		operator.Resources.IAM.GetUser.WithUserName("zhanghaonan"))

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Header["X-Request-Id"])
	fmt.Println(string(body))
}
