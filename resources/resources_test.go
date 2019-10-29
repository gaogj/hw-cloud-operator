package Api

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestResourcesNew(t *testing.T) {
	Resources := NewResources("4e7b816ca***ba02a00458")
	resources,err := Resources.VPC.Get("cn-north-1.myhuaweicloud.com", Resources.VPC.Get.WithResourceId("11016528-834a-4ee6-b130-89620a674a4c"))
	if err != nil {
		fmt.Println(err)
	}

	body,err := ioutil.ReadAll(resources.Body)
	fmt.Println(string(body))
}