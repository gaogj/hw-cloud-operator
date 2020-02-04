package Api

type Resources struct {
	VPC *VPC
	IAM *IAM
	ECS *ECS
	EVS *EVS
	BMS *BMS
	SCM *SCM
	DCS *DCS
	RDS *RDS
	ELB *ELB
	//DMS *DMS
}

type VPC struct {
	GetVPC VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

type ECS struct {
	GetECS ECSGet
}

type EVS struct {
	GetEVS EVSGet
}

type BMS struct {
	GetBMS BMSGet
}

type SCM struct {
	ListCerts CertsList
}

type DCS struct {
	GetDCS DCSGet
}

type RDS struct {
	GetRDS RDSGet
}

type ELB struct {
	GetELB ELBGet
	GetPool PoolGet
	ListPool PoolList
	GetListener ListenerGet
	GetHealtHmonitor HealthMonitorGet
	ListHealtHmonitor HealthMonitorList
	GetL7Policies L7PolicieGet
	ListL7Policies L7PolicieList
	ListL7PolicieRules L7PolicieRulesList
}

type IAM struct {
	GetUser UserGet
	GetGroup GroupGet
	CreateUser UserCreate
	AddUserToGroup GroupAddUser
	ListGroupsForUser ListGroupsForUser
	//GetUserByGroup UserByGroupGet
}

//type DMS struct {
//	CreatingConsumer
//}

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
			GetECS: newECSGetFunc(),
		},
		EVS: &EVS{
			GetEVS: newEVSGetFunc(),
		},
		BMS: &BMS{
			GetBMS: newBMSGetFunc(),
		},
		SCM: &SCM{
			ListCerts: newCertsListFunc(),
		},
		DCS: &DCS{
			GetDCS: newDCSGetFunc(),
		},
		RDS: &RDS{
			GetRDS: newRDSGetFunc(),
		},
		ELB: &ELB{
			GetELB: newELBGetFunc(),
			GetPool: newPoolGetFunc(),
			ListPool: newPoolListFunc(),
			GetListener: newListenerGetFunc(),
			GetHealtHmonitor: newHealthMonitorGetFunc(),
			ListHealtHmonitor: newHealthMonitorListFunc(),
			GetL7Policies: newL7PolicieGetFunc(),
			ListL7Policies: newL7PolicieListFunc(),
			ListL7PolicieRules: newL7PolicieRulesListFunc(),
		},
	}
}
