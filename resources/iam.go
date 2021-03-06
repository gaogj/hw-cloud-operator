package Api

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return func(endpoint string, name string, o ...func(*GroupGetRequest)) (*http.Response, error) {
		var r = GroupGetRequest{
			GroupName: name,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type GroupGet func(endpoint string, name string, o ...func(*GroupGetRequest)) (*http.Response, error)

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

//CreateUserInfo
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

//AddUserToGroup
func newGroupAddUserFunc() GroupAddUser {
	return func(endpoint, groupid, userid string, o ...func(*GroupAddUserRequest)) (*http.Response, error) {
		var r = GroupAddUserRequest{
			GroupId: groupid,
			UserId: userid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type GroupAddUser func(endpoint, groupid, userid string, o ...func(*GroupAddUserRequest)) (*http.Response, error)

type GroupAddUserRequest struct {
	ProjectId string
	Endpoint string

	GroupId string
	UserId string
}

func (gau GroupAddUserRequest) Do() (*http.Response, error) {
	if Endpoints[gau.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if gau.UserId == "" {
		return nil,errors.New("UserId cannot be empty")
	}
	if gau.GroupId == "" {
		return nil,errors.New("GroupId cannot be empty")
	}

	RequestInfo := RequestInfo{
		resourceId: fmt.Sprintf("%s/users/%s", gau.GroupId,gau.UserId),
		endpoint:Endpoints[gau.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "groups",
		method: "PUT",
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

func newListGroupsForUserFunc() ListGroupsForUser {
	return func(endpoint, userid string, o ...func(*ListGroupsForUserRequest)) (*http.Response, error) {
		var r = ListGroupsForUserRequest{
			UserId: userid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type ListGroupsForUser func(endpoint, userid string, o ...func(*ListGroupsForUserRequest)) (*http.Response, error)

type ListGroupsForUserRequest struct {
	ProjectId string
	Endpoint string

	UserId string
}

func (lgfu ListGroupsForUserRequest) Do() (*http.Response, error) {
	if Endpoints[lgfu.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if lgfu.UserId == "" {
		return nil,errors.New("UserId cannot be empty")
	}

	RequestInfo := RequestInfo{
		resourceId: fmt.Sprintf("%s/groups",lgfu.UserId),
		endpoint:Endpoints[lgfu.Endpoint].Host,
		apiVersion: "v3",
		category: "iam",
		apiObject: "users",
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

//DeleteUserInfo
func newUserDeleteFunc() UserDelete {
	return func(endpoint, userId string, o ...func(*UserDeleteRequest)) (*http.Response, error) {
		var r = UserDeleteRequest{
			UserId: userId,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type UserDelete func(endpoint, userId string, o ...func(*UserDeleteRequest)) (*http.Response, error)

type UserDeleteRequest struct {
	ProjectId string
	Endpoint string

	UserId string
}

func (ud UserDeleteRequest) Do() (*http.Response, error) {
	var body *bytes.Buffer

	body = new(bytes.Buffer)
	defer body.Reset()

	if Endpoints[ud.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		endpoint:Endpoints[ud.Endpoint].Host,
		resourceId: ud.UserId,
		apiVersion: "v3",
		category: "iam",
		apiObject: "users",
		method: "DELETE",
		scheme: "https",
		body: nil,
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
