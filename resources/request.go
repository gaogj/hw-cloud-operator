package Api

import (
	"github.com/gaogj/hw-cloud-operator/utils"
	"net"
	"net/http"
	"net/url"
	"time"
)

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

func newRequest(method string, url *url.URL, params ...map[string]string) (*http.Request, error) {
	r := &http.Request{
		Method:     method,
		URL:        url,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	if len(params) > 0 {
		q := r.URL.Query()
		for _, v := range params {
			for k, vv := range v {
				q.Set(k, vv)
			}
		}
		r.URL.RawQuery = q.Encode()
	}

	r.Header.Add("content-type", "application/json")
	Sign.Sign(r)

	return r, nil
}