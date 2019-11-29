package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//UserGet
func newEcsGetFunc() EcsGet {
	return func(endpoint, serverid string, o ...func(*EcsGetRequest)) (*http.Response, error) {
		var r = EcsGetRequest{
			ServerId: serverid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EcsGet func(endpoint, serverid string, o ...func(*EcsGetRequest)) (*http.Response, error)

type EcsGetRequest struct {
	ProjectId string
	Endpoint string

	ServerId string
}

func (eg EcsGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
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
