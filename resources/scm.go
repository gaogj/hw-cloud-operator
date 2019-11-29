package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//UserGet
func newSCMListFunc() SCMList {
	return func(endpoint, orderStatus string, o ...func(*SCMListRequest)) (*http.Response, error) {
		var r = SCMListRequest{
			OrderStatus: orderStatus,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type SCMList func(endpoint, orderStatus string, o ...func(*SCMListRequest)) (*http.Response, error)

type SCMListRequest struct {
	ProjectId string
	Endpoint string

	OrderStatus string
	Marker string
	Limit string
}

func (sl SCMListRequest) Do() (*http.Response, error) {
	if Endpoints[sl.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[sl.Endpoint].ProjectId,
		endpoint:Endpoints[sl.Endpoint].Host,
		apiVersion: "v2",
		category: "scm",
		apiObject: "scm/certlist",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}
	//params
	if sl.OrderStatus != "" {
		RequestInfo.params["order_status"] = sl.OrderStatus
	}

	if sl.Marker != "" {
		RequestInfo.params["offset"] = sl.Marker
	}

	if sl.Limit != "" {
		RequestInfo.params["limit"] = sl.Limit
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

func (sl SCMList) WithOrderStatus(OrderStatus string) func(*SCMListRequest) {
	return func(slr *SCMListRequest) {
		slr.OrderStatus = OrderStatus
	}
}

func (sl SCMList) WithMarker(Marker string) func(*SCMListRequest) {
	return func(slr *SCMListRequest) {
		slr.Marker = Marker
	}
}

func (sl SCMList) WithLimit(Limit string) func(*SCMListRequest) {
	return func(slr *SCMListRequest) {
		slr.Limit = Limit
	}
}
