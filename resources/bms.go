package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//BMSGet
func newBMSGetFunc() BMSGet {
	return func(endpoint, serverid string, o ...func(*BMSGetRequest)) (*http.Response, error) {
		var r = BMSGetRequest{
			ServerId: serverid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type BMSGet func(endpoint, serverid string, o ...func(*BMSGetRequest)) (*http.Response, error)

type BMSGetRequest struct {
	ProjectId string
	Endpoint string

	ServerId string
}

func (eg BMSGetRequest) Do() (*http.Response, error) {
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
		category: "bms",
		apiObject: "baremetalservers",
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

//BMSList
func newBMSListFunc() BMSList {
	return func(endpoint, offset, limit string, o ...func(*BMSListRequest)) (*http.Response, error) {
		var r = BMSListRequest{
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

type BMSList func(endpoint, offset, limit string, o ...func(*BMSListRequest)) (*http.Response, error)

type BMSListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el BMSListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v1",
		category: "bms",
		apiObject: "baremetalservers/detail",
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