package Api

import (
	"github.com/pkg/errors"
	"net/http"
)

//EcsGet
func newELBGetFunc() ELBGet {
	return func(endpoint string, o ...func(*ELBGetRequest)) (*http.Response, error) {
		var r = ELBGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type ELBGet func(endpoint string, o ...func(*ELBGetRequest)) (*http.Response, error)

type ELBGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string
}

func (eg ELBGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v1.0",
		category: "elb",
		apiObject: "elbaas/loadbalancers",
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

func (eg ELBGet) WithResourceId(ResourceId string) func(*ELBGetRequest) {
	return func(egr *ELBGetRequest) {
		egr.ResourceId = ResourceId
	}
}

//listeners
func newListenerGetFunc() ListenerGet {
	return func(endpoint string, o ...func(*ListenerGetRequest)) (*http.Response, error) {
		var r = ListenerGetRequest{
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type ListenerGet func(endpoint string, o ...func(*ListenerGetRequest)) (*http.Response, error)

type ListenerGetRequest struct {
	ProjectId string
	Endpoint string

	ResourceId string
}

func (eg ListenerGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.ResourceId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v1.0",
		category: "elb",
		apiObject: "elbaas/listeners",
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

func (lg ListenerGet) WithResourceId(ResourceId string) func(*ListenerGetRequest) {
	return func(lgr *ListenerGetRequest) {
		lgr.ResourceId = ResourceId
	}
}

//pools
func newPoolGetFunc() PoolGet {
	return func(endpoint, poolid string, o ...func(*PoolGetRequest)) (*http.Response, error) {
		var r = PoolGetRequest{
			PoolId: poolid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type PoolGet func(endpoint, poolid string, o ...func(*PoolGetRequest)) (*http.Response, error)

type PoolGetRequest struct {
	ProjectId string
	Endpoint string

	PoolId string
}

func (eg PoolGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.PoolId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.PoolId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/pools",
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

//list
func newPoolListFunc() PoolList {
	return func(endpoint, offset, limit string, o ...func(*PoolListRequest)) (*http.Response, error) {
		var r = PoolListRequest{
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

type PoolList func(endpoint, offset, limit string, o ...func(*PoolListRequest)) (*http.Response, error)

type PoolListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el PoolListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/pools",
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

//healthmonitors get
func newHealthMonitorGetFunc() HealthMonitorGet {
	return func(endpoint, healthmonitorid string, o ...func(*HealthMonitorGetRequest)) (*http.Response, error) {
		var r = HealthMonitorGetRequest{
			HealthmonitorId: healthmonitorid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type HealthMonitorGet func(endpoint, healthmonitorid string, o ...func(*HealthMonitorGetRequest)) (*http.Response, error)

type HealthMonitorGetRequest struct {
	ProjectId string
	Endpoint string

	HealthmonitorId string
}

func (eg HealthMonitorGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.HealthmonitorId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.HealthmonitorId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/healthmonitors",
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

//list
func newHealthMonitorListFunc() HealthMonitorList {
	return func(endpoint, offset, limit string, o ...func(*HealthMonitorListRequest)) (*http.Response, error) {
		var r = HealthMonitorListRequest{
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

type HealthMonitorList func(endpoint, offset, limit string, o ...func(*HealthMonitorListRequest)) (*http.Response, error)

type HealthMonitorListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el HealthMonitorListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/healthmonitors",
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

//L7PolicieGet get
func newL7PolicieGetFunc() L7PolicieGet {
	return func(endpoint, l7policieid string, o ...func(*L7PolicieGetRequest)) (*http.Response, error) {
		var r = L7PolicieGetRequest{
			L7PolicieId: l7policieid,
			Endpoint: endpoint,
		}
		for _, f := range o {
			f(&r)
		}
		return r.Do()
	}
}

type L7PolicieGet func(endpoint, healthmonitorid string, o ...func(*L7PolicieGetRequest)) (*http.Response, error)

type L7PolicieGetRequest struct {
	ProjectId string
	Endpoint string

	L7PolicieId string
}

func (eg L7PolicieGetRequest) Do() (*http.Response, error) {
	if Endpoints[eg.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	if eg.L7PolicieId == "" {
		return nil,errors.New("Can't find the ecs server id")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[eg.Endpoint].ProjectId,
		resourceId: eg.L7PolicieId,
		endpoint:Endpoints[eg.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/l7policies",
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

//list
func newL7PolicieListFunc() L7PolicieList {
	return func(endpoint, offset, limit string, o ...func(*L7PolicieListRequest)) (*http.Response, error) {
		var r = L7PolicieListRequest{
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

type L7PolicieList func(endpoint, offset, limit string, o ...func(*L7PolicieListRequest)) (*http.Response, error)

type L7PolicieListRequest struct {
	ProjectId string
	Endpoint string

	Offset string
	Limit string
}

func (el L7PolicieListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/l7policies",
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

//list
func newL7PolicieRulesListFunc() L7PolicieRulesList {
	return func(endpoint, l7policyid, offset, limit string, o ...func(*L7PolicieRulesListRequest)) (*http.Response, error) {
		var r = L7PolicieRulesListRequest{
			L7PolicyId: l7policyid,
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

type L7PolicieRulesList func(endpoint, l7policyid, offset, limit string, o ...func(*L7PolicieRulesListRequest)) (*http.Response, error)

type L7PolicieRulesListRequest struct {
	ProjectId string
	Endpoint string

	L7PolicyId string

	Offset string
	Limit string
}

func (el L7PolicieRulesListRequest) Do() (*http.Response, error) {
	if Endpoints[el.Endpoint].Host == "" {
		return nil,errors.New("Can't find the Endpoint host")
	}

	RequestInfo := RequestInfo{
		projectId: Endpoints[el.Endpoint].ProjectId,
		resourceId: el.L7PolicyId + "/rules",
		endpoint:Endpoints[el.Endpoint].Host,
		apiVersion: "v2.0",
		category: "elb",
		apiObject: "lbaas/l7policies",
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
