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

func (ur UserGetRequest) Do() (*http.Response, error) {

	if Endpoints[ur.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		endpoint:Endpoints[ur.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "users",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if ur.UserName != "" {
		RequestInfo.params["name"] = ur.UserName
	}

	if ur.Enabled != "" {
		RequestInfo.params["enabled"] = ur.Enabled
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
	return func(ur *UserGetRequest) {
		ur.UserName = UserName
	}
}

func (vg UserGet) WithEnabled(Enabled string) func(*UserGetRequest) {
	return func(ur *UserGetRequest) {
		ur.Enabled = Enabled
	}
}

//GetUserInfo
func newUserCreateFunc() UserCreate {
	return func(endpoint, name, password string, o ...func(*UserCreateRequest)) (*http.Response, error) {
		var r = UserCreateRequest{
			Name: name,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type UserCreate func(endpoint, name, password string, o ...func(*UserCreateRequest)) (*http.Response, error)

type UserCreateRequest struct {
	ProjectId string
	Endpoint string

	Name string
	Password string

	Description string
}

func (uc UserCreateRequest) Do() (*http.Response, error) {

	if Endpoints[uc.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		resourceId: uc.Name,
		endpoint:Endpoints[uc.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "/users",
		method: "POST",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	RequestInfo.params["name"] = uc.Name
	RequestInfo.params["password"] = uc.Password

	if uc.Description != "" {
		RequestInfo.params["description"] = uc.Description
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

func (vg UserGet) WithDescription(description string) func(*UserCreateRequest) {
	return func(uc *UserCreateRequest) {
		uc.Description = description
	}
}