package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newDCSGetFunc() DCSGet {
	return func(endpoint, instanceid string, o ...func(*DCSGetRequest)) (*http.Response, error) {
		var r = DCSGetRequest{
			InstanceId: instanceid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type DCSGet func(endpoint, instanceid string, o ...func(*DCSGetRequest)) (*http.Response, error)

type DCSGetRequest struct {
	ProjectId string
	Endpoint string

	InstanceId string
}

func (eg DCSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.InstanceId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.InstanceId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v1.0",
		category: "dcs",
		apiObject: "instances",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

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
func newDCSListFunc() DCSList {
	return func(endpoint, offset, limit string, o ...func(*DCSListRequest)) (*http.Response, error) {
		var r = DCSListRequest{
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

type DCSList func(endpoint, offset, limit string, o ...func(*DCSListRequest)) (*http.Response, error)

type DCSListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el DCSListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v1.0",
		category: "dcs",
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