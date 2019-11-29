package Api

type Resources struct {
	VPC *VPC
	IAM *IAM
	ECS *ECS
	SCM *SCM
}

type VPC struct {
	GetVPC VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

type ECS struct {
	GetECS EcsGet
}

type SCM struct {
	ListCerts CertsList
}

type IAM struct {
	GetUser UserGet
	GetGroup GroupGet
	CreateUser UserCreate
	AddUserToGroup GroupAddUser
	ListGroupsForUser ListGroupsForUser
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
			AddUserToGroup: newGroupAddUserFunc(),
			ListGroupsForUser: newListGroupsForUserFunc(),
		},
		ECS: &ECS{
			GetECS: newEcsGetFunc(),
		},
		SCM: &SCM{
			ListCerts: newCertsListFunc(),
		},
	}
}
