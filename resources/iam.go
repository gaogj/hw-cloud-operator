package Api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

//UserGet
func newUserGetFunc(ProjectId string) UserGet {
	return func(endpoint string, o ...func(*UserGetRequest)) (*http.Response, error) {
		var r = UserGetRequest{
			ProjectId: ProjectId,
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
	var (
		method string
		scheme string
		apiVersion string
		apiObject string
		params map[string]string
		path   strings.Builder
		category string  = "iam"
	)

	apiVersion = "v3"
	apiObject = "users"
	method = "GET"
	scheme = "https"
	params = make(map[string]string)

	if vr.Endpoint == "" {
		return nil,errors.New("Can't find the Endpoint")
	}
	vr.Endpoint = category + "." + vr.Endpoint

	path.WriteString("/")
	path.WriteString(apiVersion)

	path.WriteString("/")
	path.WriteString(apiObject)

	//if vr.ResourceId != "" {
	//	path.WriteString("/")
	//	path.WriteString(vr.ResourceId)
	//}

	url := &url.URL{
		Scheme: scheme,
		Host: vr.Endpoint,
		Path: path.String(),
	}

	//params
	if vr.UserName != "" {
		params["UserName"] = vr.UserName
	}

	if vr.Enabled != "" {
		params["Enabled"] = vr.Enabled
	}

	req, _ := newRequest(method, url, params)

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
