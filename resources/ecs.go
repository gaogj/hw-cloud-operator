package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newECSGetFunc() ECSGet {
	return func(endpoint string, o ...func(*ECSGetRequest)) (*http.Response, error) {
		var r = ECSGetRequest{
			ResourceId: "detail",
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type ECSGet func(endpoint string, o ...func(*ECSGetRequest)) (*http.Response, error)

type ECSGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string

	Limit string
	Offset string
}

func (eg ECSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.ResourceId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v1",
		category: "ecs",
		apiObject: "cloudservers",
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

func (eg ECSGet) WithResourceId(ResourceId string) func(*ECSGetRequest) {
	return func(egr *ECSGetRequest) {
		egr.ResourceId = ResourceId
	}
}

func (eg ECSGet) WithMarker(Marker string) func(*ECSGetRequest) {
	return func(egr *ECSGetRequest) {
		egr.Offset = Marker
	}
}

func (eg ECSGet) WithLimit(Limit string) func(*ECSGetRequest) {
	return func(egr *ECSGetRequest) {
		egr.Limit = Limit
	}
}
