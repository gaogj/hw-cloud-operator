package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newRDSGetFunc() RDSGet {
	return func(endpoint, instanceid string, o ...func(*RDSGetRequest)) (*http.Response, error) {
		var r = RDSGetRequest{
			InstanceId: instanceid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type RDSGet func(endpoint, serverid string, o ...func(*RDSGetRequest)) (*http.Response, error)

type RDSGetRequest struct {
	ProjectId string
	Endpoint string

	InstanceId string
}

func (eg RDSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.InstanceId == "" {
		return nil,errors.New("Can't find the ecs instance id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		//resourceId: eg.InstanceId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v3",
		category: "rds",
		apiObject: "instances",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	RequestInfo.params["id"] = eg.InstanceId

	req, err := newRequest(RequestInfo)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//EcsList
func newRDSListFunc() RDSList {
	return func(endpoint, offset, limit string, o ...func(*RDSListRequest)) (*http.Response, error) {
		var r = RDSListRequest{
			Endpoint: endpoint,
			Offset: offset,
			Limit: limit,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type RDSList func(endpoint, offset, limit string, o ...func(*RDSListRequest)) (*http.Response, error)

type RDSListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el RDSListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v3",
		category: "rds",
		apiObject: "instances",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	RequestInfo.params["offset"] = el.Offset
	RequestInfo.params["limit"] = el.Limit

	req, err := newRequest(RequestInfo)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}