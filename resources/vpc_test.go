package Api

import (
	"fmt"
	"testing"
)

func TestVPCGetRequest(t *testing.T) {
	r := VPCGetRequest{
		ProjectId: "4e7b816ca65******00dba02a00458",
		Endpoint: "cn-north-1.myhuaweicloud.com",
	}
	fmt.Println(r.Do())
}