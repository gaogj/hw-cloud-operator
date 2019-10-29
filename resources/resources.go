package Api

type Resources struct {
	VPC *VPC
}

type VPC struct {
	Get VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

func NewResources(ProjectId string) *Resources {
	return &Resources{
		VPC: &VPC{
			Get: newVPCGetFunc(ProjectId),
			GetSubnet: newSubnetGetFunc(ProjectId),
		},
	}
}
