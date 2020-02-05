package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newSLBGetFunc() SLBGet {
	return func(endpoint string, o ...func(*SLBGetRequest)) (*http.Response, error) {
		var r = SLBGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SLBGet func(endpoint string, o ...func(*SLBGetRequest)) (*http.Response, error)

type SLBGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string
}

func (eg SLBGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		endpoint: Endpoints[eg.Endpoint].Host,
		apiVersion: "v2",
		category: "elb",
		apiObject: "elb/loadbalancers",
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

func (eg SLBGet) WithResourceID(ResourceId string) func(*SLBGetRequest) {
	return func(egr *SLBGetRequest) {
		egr.ResourceId = ResourceId
	}
}
