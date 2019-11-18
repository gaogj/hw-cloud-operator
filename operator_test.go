package hwcloudoperator

import (
	"fmt"
	"github.com/gaogj/hw-cloud-operator/utils"
	"io/ioutil"
	"testing"
)

//
func TestVPCNewOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
<<<<<<< HEAD
		AccessKey: "FGEXLWUDPS9OJHD2DP30",
		SecretAccessKey: "t3yIrYtyn9KZPJlCIvt0e6Dw729OXxJeUWODMjzo",
		AccountId: "a2f58a7bbf264053ab03375e7dbf9501",
=======
		AccessKey: "********",
		SecretAccessKey: "********",
		AccountId: "********",
>>>>>>> a9e0edfa8e68ac20dafa1c7ef41e013e6c3798a1
		Endpoints: 	map[string]utils.Endpoint{
			"bj1": utils.Endpoint{
				Host: "cn-north-1.myhuaweicloud.com",
				ProjectId: "********",
			},
		},
	})

	res,err := operator.Resources.VPC.GetVPC("bj1",
		operator.Resources.VPC.GetVPC.WithResourceId("11016528-834a-4ee6-b130-********"))

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
////
func TestNewSubnetOperator(t *testing.T) {
	operator := NewOperator(&utils.Config{
<<<<<<< HEAD
		AccessKey: "FGEXLWUDPS9OJHD2DP30",
		SecretAccessKey: "t3yIrYtyn9KZPJlCIvt0e6Dw729OXxJeUWODMjzo",
		AccountId: "a2f58a7bbf264053ab03375e7dbf9501",
=======
		AccessKey: "********",
		SecretAccessKey: "********",
		AccountId: "********",
>>>>>>> a9e0edfa8e68ac20dafa1c7ef41e013e6c3798a1
		Endpoints: 	map[string]utils.Endpoint{
			"bj1": utils.Endpoint{
				Host: "cn-north-1.myhuaweicloud.com",
				ProjectId: "********",
			},
		},
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
<<<<<<< HEAD
		AccessKey: "FGEXLWUDPS9OJHD2DP30",
		SecretAccessKey: "t3yIrYtyn9KZPJlCIvt0e6Dw729OXxJeUWODMjzo",
		AccountId: "a2f58a7bbf264053ab03375e7dbf9501",
=======
		AccessKey: "********",
		SecretAccessKey: "********",
		AccountId: "********",
>>>>>>> a9e0edfa8e68ac20dafa1c7ef41e013e6c3798a1
		Endpoints: 	map[string]utils.Endpoint{
			"bj1": utils.Endpoint{
				Host: "cn-north-1.myhuaweicloud.com",
				ProjectId: "********",
			},
		},
	})

	res,err := operator.Resources.IAM.GetUser("bj1")

	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(res.Header["X-Request-Id"])
	fmt.Println(string(body))
}
