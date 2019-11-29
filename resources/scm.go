package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//UserGet
func newCertsListFunc() CertsList {
	return func(endpoint, orderStatus string, o ...func(*CertsListRequest)) (*http.Response, error) {
		var r = CertsListRequest{
			OrderStatus: orderStatus,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type CertsList func(endpoint, orderStatus string, o ...func(*CertsListRequest)) (*http.Response, error)

type CertsListRequest struct {
	ProjectId string
	Endpoint string

	OrderStatus string
	Marker string
	Limit string
}

func (sl CertsListRequest) Do() (*http.Response, error) {
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

func (sl CertsList) WithOrderStatus(OrderStatus string) func(*CertsListRequest) {
	return func(slr *CertsListRequest) {
		slr.OrderStatus = OrderStatus
	}
}

func (sl CertsList) WithMarker(Marker string) func(*CertsListRequest) {
	return func(slr *CertsListRequest) {
		slr.Marker = Marker
	}
}

func (sl CertsList) WithLimit(Limit string) func(*CertsListRequest) {
	return func(slr *CertsListRequest) {
		slr.Limit = Limit
	}
}
