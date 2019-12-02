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
}

type VPC struct {
	GetVPC VPCGet
	GetSubnet SubnetGet
	GetPublicip PublicipGet
}

type ECS struct {
	GetECS ECSGet
	ListECS ECSList
}

type EVS struct {
	GetEVS EVSGet
	ListEVS EVSList
}

type BMS struct {
	GetBMS BMSGet
	ListBMS BMSList
}

type SCM struct {
	ListCerts CertsList
}

type DCS struct {
	GetDCS DCSGet
	ListDCS DCSList
}

type RDS struct {
	GetRDS RDSGet
	ListRDS RDSList
}

type ELB struct {
	GetELB ELBGet
	ListELB ELBList
	GetPool PoolGet
	ListPool PoolList
	GetListener ListenerGet
	ListListener ListenerList
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
			ListECS: newECSListFunc(),
		},
		EVS: &EVS{
			GetEVS: newEVSGetFunc(),
			ListEVS: newEVSListFunc(),
		},
		BMS: &BMS{
			GetBMS: newBMSGetFunc(),
			ListBMS: newBMSListFunc(),
		},
		SCM: &SCM{
			ListCerts: newCertsListFunc(),
		},
		DCS: &DCS{
			GetDCS: newDCSGetFunc(),
			ListDCS: newDCSListFunc(),
		},
		RDS: &RDS{
			GetRDS: newRDSGetFunc(),
			ListRDS: newRDSListFunc(),
		},
		ELB: &ELB{
			GetELB: newELBGetFunc(),
			ListELB: newELBListFunc(),
			GetPool: newPoolGetFunc(),
			ListPool: newPoolListFunc(),
			GetListener: newListenerGetFunc(),
			ListListener: newListenerListFunc(),
			GetHealtHmonitor: newHealthMonitorGetFunc(),
			ListHealtHmonitor: newHealthMonitorListFunc(),
			GetL7Policies: newL7PolicieGetFunc(),
			ListL7Policies: newL7PolicieListFunc(),
			ListL7PolicieRules: newL7PolicieRulesListFunc(),
		},
	}
}
