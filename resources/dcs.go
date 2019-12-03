package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newDCSGetFunc() DCSGet {
	return func(endpoint string, o ...func(*DCSGetRequest)) (*http.Response, error) {
		var r = DCSGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type DCSGet func(endpoint string, o ...func(*DCSGetRequest)) (*http.Response, error)

type DCSGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string

	Limit string
	Offset string
}

func (dg DCSGetRequest) Do() (*http.Response, error) {
	if Endpoints[dg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[dg.Endpoint].ProjectId,
		endpoint:Endpoints[dg.Endpoint].Host,
		apiVersion: "v1.0",
		category: "dcs",
		apiObject: "instances",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if dg.Offset != "" {
		RequestInfo.params["marker"] = dg.Offset
	}

	if dg.Limit != "" {
		RequestInfo.params["limit"] = dg.Limit
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

func (dg DCSGet) WithResourceId(ResourceId string) func(*DCSGetRequest) {
	return func(dgr *DCSGetRequest) {
		dgr.ResourceId = ResourceId
	}
}

func (dg DCSGet) WithMarker(Marker string) func(*DCSGetRequest) {
	return func(dgr *DCSGetRequest) {
		dgr.Offset = Marker
	}
}

func (dg DCSGet) WithLimit(Limit string) func(*DCSGetRequest) {
	return func(dgr *DCSGetRequest) {
		dgr.Limit = Limit
	}
}
