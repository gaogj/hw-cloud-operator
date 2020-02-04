package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newEVSGetFunc() EVSGet {
	return func(endpoint string, o ...func(*EVSGetRequest)) (*http.Response, error) {
		var r = EVSGetRequest{
			ResourceId: "detail",
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EVSGet func(endpoint string, o ...func(*EVSGetRequest)) (*http.Response, error)

type EVSGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string

	Offset string
	Limit string
}

func (eg EVSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.ResourceId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v3",
		category: "evs",
		apiObject: "os-vendor-volumes",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
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

func (eg EVSGet) WithResourceID(ResourceId string) func(*EVSGetRequest) {
	return func(egr *EVSGetRequest) {
		egr.ResourceId = ResourceId
	}
}

func (eg EVSGet) WithMarker(Marker string) func(*EVSGetRequest) {
	return func(egr *EVSGetRequest) {
		egr.Offset = Marker
	}
}

func (eg EVSGet) WithLimit(Limit string) func(*EVSGetRequest) {
	return func(egr *EVSGetRequest) {
		egr.Limit = Limit
	}
}
