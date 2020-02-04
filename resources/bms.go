package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//BMSGet
func newBMSGetFunc() BMSGet {
	return func(endpoint string, o ...func(*BMSGetRequest)) (*http.Response, error) {
		var r = BMSGetRequest{
			ResourceId: "detail",
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type BMSGet func(endpoint string, o ...func(*BMSGetRequest)) (*http.Response, error)

type BMSGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string

	Limit string
	Offset string
}

func (bg BMSGetRequest) Do() (*http.Response, error) {
	if Endpoints[bg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[bg.Endpoint].ProjectId,
		resourceId: bg.ResourceId,
		endpoint: Endpoints[bg.Endpoint].Host,
		apiVersion: "v1",
		category: "bms",
		apiObject: "baremetalservers",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if bg.Offset != "" {
		RequestInfo.params["marker"] = bg.Offset
	}

	if bg.Limit != "" {
		RequestInfo.params["limit"] = bg.Limit
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

func (bg BMSGet) WithResourceID(ResourceId string) func(*BMSGetRequest) {
	return func(bgr *BMSGetRequest) {
		bgr.ResourceId = ResourceId
	}
}

func (bg BMSGet) WithMarker(Marker string) func(*BMSGetRequest) {
	return func(bgr *BMSGetRequest) {
		bgr.Offset = Marker
	}
}

func (bg BMSGet) WithLimit(Limit string) func(*BMSGetRequest) {
	return func(bgr *BMSGetRequest) {
		bgr.Limit = Limit
	}
}
