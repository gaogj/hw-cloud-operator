package Api

import (
	//"github.com/gaogj/hw-cloud-operator/utils"
	"net/http"
	"github.com/pkg/errors"
)

//UserGet
func newUserGetFunc() UserGet {
	return func(endpoint string, o ...func(*UserGetRequest)) (*http.Response, error) {
		var r = UserGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type UserGet func(endpoint string, o ...func(*UserGetRequest)) (*http.Response, error)

type UserGetRequest struct {
	ProjectId string
	Endpoint string

	UserName string
	Enabled string
}

func (vr UserGetRequest) Do() (*http.Response, error) {

	if Endpoints[vr.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		endpoint:Endpoints[vr.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "users",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if vr.UserName != "" {
		RequestInfo.params["name"] = vr.UserName
	}

	if vr.Enabled != "" {
		RequestInfo.params["enabled"] = vr.Enabled
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

func (vg UserGet) WithUserName(UserName string) func(*UserGetRequest) {
	return func(vr *UserGetRequest) {
		vr.UserName = UserName
	}
}

func (vg UserGet) WithEnabled(Enabled string) func(*UserGetRequest) {
	return func(vr *UserGetRequest) {
		vr.Enabled = Enabled
	}
}
