package Api

import (
	"github.com/gaogj/hw-cloud-operator/utils"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestInfo struct {
	category string
	projectId string
	endpoint string
	method string
	scheme string
	apiVersion string
	apiObject string
	params map[string]string
	resourceId string
}

var (
	httpClient *http.Client
	Sign *utils.Signer
)

func InitHttpClient(config *utils.Config) {
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
		Key:    config.AccessKey,
		Secret: config.SecretAccessKey,
	}
}

func newRequest(RequestInfo RequestInfo) (*http.Request, error) {
	var path strings.Builder

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
			//for k, vv := range v {
			//	q.Set(k, vv)
			//}
		}
		r.URL.RawQuery = q.Encode()
	}

	r.Header.Add("content-type", "application/json")

	if RequestInfo.projectId == "" {
		r.Header.Add("X-Domain-Id", "a2f58a7bbf264053ab03375e7dbf9501")
	}

	Sign.Sign(r)

	return r, nil
}