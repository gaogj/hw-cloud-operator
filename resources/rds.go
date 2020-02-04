package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newRDSGetFunc() RDSGet {
	return func(endpoint string, o ...func(*RDSGetRequest)) (*http.Response, error) {
		var r = RDSGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type RDSGet func(endpoint string, o ...func(*RDSGetRequest)) (*http.Response, error)

type RDSGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string

	Offset string
	Limit string
}

func (eg RDSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v3",
		category: "rds",
		apiObject: "instances",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if eg.ResourceId != "" {
		RequestInfo.params["id"] = eg.ResourceId
	}

	if eg.Offset != "" {
		RequestInfo.params["marker"] = eg.Offset
	}

	if eg.Limit != "" {
		RequestInfo.params["limit"] = eg.Limit
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

func (dg RDSGet) WithResourceID(ResourceId string) func(*RDSGetRequest) {
	return func(dgr *RDSGetRequest) {
		dgr.ResourceId = ResourceId
	}
}

func (dg RDSGet) WithMarker(Marker string) func(*RDSGetRequest) {
	return func(dgr *RDSGetRequest) {
		dgr.Offset = Marker
	}
}

func (dg RDSGet) WithLimit(Limit string) func(*RDSGetRequest) {
	return func(dgr *RDSGetRequest) {
		dgr.Limit = Limit
	}
}
