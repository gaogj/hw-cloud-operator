package Api

import (
	"bytes"
	"github.com/gaogj/hw-cloud-operator/utils"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	httpClient *http.Client
	Sign *utils.Signer
	Endpoints map[string]utils.Endpoint
)

type RequestInfo struct {
	category string
	projectId string
	endpoint string
	method string
	scheme string
	apiVersion string
	apiObject string
	body io.Reader
	params map[string]string
	resourceId string
}

func InitHttpClient(cfg *utils.Config) {
	Endpoints = cfg.Endpoints
	httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   30,
			ResponseHeaderTimeout: time.Second*120,
			DialContext:           (&net.Dialer{
				Timeout: 30 * time.Second,
				KeepAlive: 30 * time.Second}).DialContext,
		},
	}

	Sign = &utils.Signer{
		Key:    cfg.AccessKey,
		Secret: cfg.SecretAccessKey,
		Id:     cfg.AccountId,
	}
}

func newRequest(RequestInfo RequestInfo) (*http.Request, error) {
	var path strings.Builder

	if RequestInfo.endpoint == "" {
		return nil, errors.New("endpoint can't be empty")
	}

	if RequestInfo.category == "" {
		return nil, errors.New("category can't be empty")
	}

	if RequestInfo.method == "" {
		return nil, errors.New("method can't be empty")
	}

	if RequestInfo.scheme == "" {
		return nil, errors.New("scheme can't be empty")
	}

	if RequestInfo.apiVersion == "" {
		return nil, errors.New("apiVersion can't be empty")
	}
	path.WriteString("/")
	path.WriteString(RequestInfo.apiVersion)

	if RequestInfo.projectId != "" {
		path.WriteString("/")
		path.WriteString(RequestInfo.projectId)
	}

	if RequestInfo.apiObject == "" {
		return nil, errors.New("apiObject can't be empty")
	}
	path.WriteString("/")
	path.WriteString(RequestInfo.apiObject)

	if RequestInfo.resourceId != "" {
		path.WriteString("/")
		path.WriteString(RequestInfo.resourceId)
	}

	apiEndpoint := RequestInfo.category +  "." + RequestInfo.endpoint

	url := &url.URL{
		Scheme: RequestInfo.scheme,
		Host: apiEndpoint,
		Path: path.String(),
	}

	r := &http.Request{
		Method:     RequestInfo.method,
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if len(RequestInfo.params) > 0 {
		q := r.URL.Query()
		for k, v := range RequestInfo.params {
			q.Set(k, v)
		}
		r.URL.RawQuery = q.Encode()
	}

	if RequestInfo.body != nil {
		switch b := RequestInfo.body.(type) {
		case *bytes.Buffer:
			r.Body = ioutil.NopCloser(RequestInfo.body)
			r.ContentLength = int64(b.Len())
		case *bytes.Reader:
			r.Body = ioutil.NopCloser(RequestInfo.body)
			r.ContentLength = int64(b.Len())
		case *strings.Reader:
			r.Body = ioutil.NopCloser(RequestInfo.body)
			r.ContentLength = int64(b.Len())
		default:
			r.Body = ioutil.NopCloser(RequestInfo.body)
		}
	}

	r.Header.Add("content-type", "application/json")

	if RequestInfo.projectId == "" {
		r.Header.Add("X-Domain-Id", Sign.Id)
	}

	Sign.Sign(r)

	return r, nil
}