package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

const category string  = "vpc"

//VPCGet
func newVPCGetFunc(ProjectId string) VPCGet {
	return func(endpoint string, o ...func(*VPCGetRequest)) (*http.Response, error) {
		var r = VPCGetRequest{
			ProjectId: ProjectId,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type VPCGet func(endpoint string, o ...func(*VPCGetRequest)) (*http.Response, error)

type VPCGetRequest struct {
	ProjectId string
	Endpoint string

	Marker string
	Limit string
	ResourceId string
}

func (vr VPCGetRequest) Do() (*http.Response, error) {
	if vr.ProjectId == "" {
		return nil,errors.New("Can't find the Endpoint")
	}

	RequestInfo := RequestInfo{
		projectId: vr.ProjectId,
		endpoint: vr.Endpoint,
		apiVersion: "v1",
		category: "vpc",
		apiObject: "vpcs",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if vr.Marker != "" {
		RequestInfo.params["marker"] = vr.Marker
	}

	if vr.Limit != "" {
		RequestInfo.params["limit"] = vr.Limit
	}

	req, _ := newRequest(RequestInfo)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vg VPCGet) WithResourceId(ResourceId string) func(*VPCGetRequest) {
	return func(vr *VPCGetRequest) {
		vr.ResourceId = ResourceId
	}
}

func (vg VPCGet) WithMarker(Marker string) func(*VPCGetRequest) {
	return func(vr *VPCGetRequest) {
		vr.Marker = Marker
	}
}

func (vg VPCGet) WithLimit(Limit string) func(*VPCGetRequest) {
	return func(vr *VPCGetRequest) {
		vr.Limit = Limit
	}
}

//SubnetGet
func newSubnetGetFunc(ProjectId string) SubnetGet {
	return func(endpoint string, o ...func(*SubnetGetRequest)) (*http.Response, error) {
		var r = SubnetGetRequest{
			ProjectId: ProjectId,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SubnetGet func(endpoint string, o ...func(*SubnetGetRequest)) (*http.Response, error)

type SubnetGetRequest struct {
	ProjectId string
	Endpoint string

	Marker string
	Limit string
	VPCId string
}

func (sr SubnetGetRequest) Do() (*http.Response, error){
	if sr.ProjectId == "" {
		return nil,errors.New("Can't find the Endpoint")
	}

	RequestInfo := RequestInfo{
		projectId: sr.ProjectId,
		endpoint: sr.Endpoint,
		apiVersion: "v1",
		category: "vpc",
		apiObject: "subnets",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if sr.Marker != "" {
		RequestInfo.params["marker"] = sr.Marker
	}

	if sr.Limit != "" {
		RequestInfo.params["limit"] = sr.Limit
	}

	req, _ := newRequest(RequestInfo)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sg SubnetGet) WithVPCId(VPCId string) func(*SubnetGetRequest) {
	return func(sr *SubnetGetRequest) {
		sr.VPCId = VPCId
	}
}

func (sg SubnetGet) WithMarker(Marker string) func(*SubnetGetRequest) {
	return func(sr *SubnetGetRequest) {
		sr.Marker = Marker
	}
}

func (sg SubnetGet) WithLimit(Limit string) func(*SubnetGetRequest) {
	return func(sr *SubnetGetRequest) {
		sr.Limit = Limit
	}
}


type PublicipGet func(endpoint string, o ...func(*PublicipGetRequest)) (*http.Response, error)

type PublicipGetRequest struct {
	Marker string
	Limit int
	ResourceId string
	VPCId string
}