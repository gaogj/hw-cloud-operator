package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newEVSGetFunc() EVSGet {
	return func(endpoint, volumeid string, o ...func(*EVSGetRequest)) (*http.Response, error) {
		var r = EVSGetRequest{
			VolumeId: volumeid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EVSGet func(endpoint, volumeid string, o ...func(*EVSGetRequest)) (*http.Response, error)

type EVSGetRequest struct {
	ProjectId string
	Endpoint string

	VolumeId string
}

func (eg EVSGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.VolumeId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.VolumeId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v3",
		category: "evs",
		apiObject: "os-vendor-volumes",
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

//EcsList
func newEVSListFunc() EVSList {
	return func(endpoint, offset, limit string, o ...func(*EVSListRequest)) (*http.Response, error) {
		var r = EVSListRequest{
			Endpoint: endpoint,
			Offset: offset,
			Limit: limit,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type EVSList func(endpoint, offset, limit string, o ...func(*EVSListRequest)) (*http.Response, error)

type EVSListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el EVSListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v3",
		category: "evs",
		apiObject: "os-vendor-volumes/detail",
		method: "GET",
		scheme: "https",
		params: make(map[string]string),
	}

	RequestInfo.params["offset"] = el.Offset
	RequestInfo.params["limit"] = el.Limit

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