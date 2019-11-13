package Api

type Resources struct {
	VPC *VPC
	IAM *IAM
}

type VPC struct {
	Get VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

type IAM struct {
	GetUser UserGet
	//GetGroup GroupGet
	//GetUserByGroup UserByGroupGet
}

func NewResources() *Resources {
	return &Resources{
		VPC: &VPC{
			Get: newVPCGetFunc(),
			GetSubnet: newSubnetGetFunc(),
		},
		IAM: &IAM{
			GetUser: newUserGetFunc(),
		},
	}
}
