package Api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
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
	var (
		method string
		scheme string
		apiVersion string
		apiObject string
		params map[string]string
		path   strings.Builder
	)

	apiVersion = "v1"
	apiObject = "vpcs"
	method = "GET"
	scheme = "https"
	params = make(map[string]string)

	if vr.Endpoint == "" {
		return nil,errors.New("Can't find the Endpoint")
	}
	vr.Endpoint = category + "." + vr.Endpoint

	path.WriteString("/")
	path.WriteString(apiVersion)

	path.WriteString("/")
	path.WriteString(vr.ProjectId)

	path.WriteString("/")
	path.WriteString(apiObject)

	if vr.ResourceId != "" {
		path.WriteString("/")
		path.WriteString(vr.ResourceId)
	}

	url := &url.URL{
		Scheme: scheme,
		Host: vr.Endpoint,
		Path: path.String(),
	}

	//params
	if vr.Marker != "" {
		params["marker"] = vr.Marker
	}

	if vr.Limit != "" {
		params["limit"] = vr.Limit
	}

	req, _ := newRequest(method, url, params)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	//body,err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

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
	var (
		method string
		scheme string
		apiVersion string
		apiObject string
		params map[string]string
		path   strings.Builder
	)

	apiVersion = "v1"
	apiObject = "subnets"
	method = "GET"
	scheme = "https"
	params = make(map[string]string)


	if sr.Endpoint == "" {
		return nil,errors.New("Can't find the Endpoint")
	}
	sr.Endpoint = category + "." + sr.Endpoint

	path.WriteString("/")
	path.WriteString(apiVersion)

	path.WriteString("/")
	path.WriteString(sr.ProjectId)

	path.WriteString("/")
	path.WriteString(apiObject)

	url := &url.URL{
		Scheme: scheme,
		Host: sr.Endpoint,
		Path: path.String(),
	}

	//params
	if sr.VPCId != "" {
		params["vpc_id"] = sr.VPCId
	}

	if sr.Marker != "" {
		params["marker"] = sr.Marker
	}

	if sr.Limit != "" {
		params["limit"] = sr.Limit
	}

	req, _ := newRequest(method, url)


	fmt.Println(req)

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	//body,err := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))

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