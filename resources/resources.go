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

func NewResources(ProjectId string) *Resources {
	return &Resources{
		VPC: &VPC{
			Get: newVPCGetFunc(ProjectId),
			GetSubnet: newSubnetGetFunc(ProjectId),
		},
		IAM: &IAM{
			GetUser: newUserGetFunc(ProjectId),
		},
	}
}
