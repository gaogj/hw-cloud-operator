package Api

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	//"github.com/gaogj/hw-cloud-operator/utils"
	"net/http"
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

func (ug UserGet) WithUserName(UserName string) func(*UserGetRequest) {
	return func(ur *UserGetRequest) {
		ur.UserName = UserName
	}
}

func (ug UserGet) WithEnabled(Enabled string) func(*UserGetRequest) {
	return func(ur *UserGetRequest) {
		ur.Enabled = Enabled
	}
}

//GetGroup
func newGroupGetFunc() GroupGet {
	return func(endpoint string, o ...func(*GroupGetRequest)) (*http.Response, error) {
		var r = GroupGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type GroupGet func(endpoint string, o ...func(*GroupGetRequest)) (*http.Response, error)

type GroupGetRequest struct {
	ProjectId string
	Endpoint string

	GroupName string
}

func (ug GroupGetRequest) Do() (*http.Response, error) {

	if Endpoints[ug.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		endpoint:Endpoints[ug.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "groups",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	//params
	if ug.GroupName != "" {
		RequestInfo.params["name"] = ug.GroupName
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

//GetUserInfo
func newUserCreateFunc() UserCreate {
	return func(endpoint, name, password string, o ...func(*UserCreateRequest)) (*http.Response, error) {
		var r = UserCreateRequest{
			Name: name,
			Password: password,
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
	var body *bytes.Buffer

	body = new(bytes.Buffer)
	defer body.Reset()

	if Endpoints[uc.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		endpoint:Endpoints[uc.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "users",
		method: "POST",
		scheme: "https",
		body: nil,
		params: make(map[string]string),
	}
	//body
	bodyString, err := json.Marshal(map[string]interface{}{
		"user": map[string]string{
			"name": uc.Name,
			"password": uc.Password,
			"description": uc.Description,
		},
	})
	if err != nil {
		return nil, err
	}

	body.Grow(len(bodyString))
	body.Write(bodyString)
	RequestInfo.body = body

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

func (uc UserCreate) WithDescription(description string) func(*UserCreateRequest) {
	return func(ur *UserCreateRequest) {
		ur.Description = description
	}
}

