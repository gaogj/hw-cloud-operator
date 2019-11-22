package Api

type Resources struct {
	VPC *VPC
	IAM *IAM
}

type VPC struct {
	GetVPC VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

type IAM struct {
	GetUser UserGet
	GetGroup GroupGet
	CreateUser UserCreate
	//GetUserByGroup UserByGroupGet
}

func NewResources() *Resources {
	return &Resources{
		VPC: &VPC{
			GetVPC: newVPCGetFunc(),
			GetSubnet: newSubnetGetFunc(),
		},
		IAM: &IAM{
			GetUser: newUserGetFunc(),
			GetGroup: newGroupGetFunc(),
			CreateUser: newUserCreateFunc(),
		},
	}
}
