package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newECSGetFunc() ECSGet {
	return func(endpoint, serverid string, o ...func(*ECSGetRequest)) (*http.Response, error) {
		var r = ECSGetRequest{
			ServerId: serverid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type ECSGet func(endpoint, serverid string, o ...func(*ECSGetRequest)) (*http.Response, error)

type ECSGetRequest struct {
	ProjectId string
	Endpoint string

	ServerId string
}

func (eg ECSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.ServerId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.ServerId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v1",
		category: "ecs",
		apiObject: "cloudservers",
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
func newECSListFunc() ECSList {
	return func(endpoint, offset, limit string, o ...func(*ECSListRequest)) (*http.Response, error) {
		var r = ECSListRequest{
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

type ECSList func(endpoint, offset, limit string, o ...func(*ECSListRequest)) (*http.Response, error)

type ECSListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el ECSListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v1",
		category: "ecs",
		apiObject: "cloudservers/detail",
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