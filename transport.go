package hms

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// needed for draining body
	respReadLimit = int64(4096)

	DefaultRetryCount      int = 5
	DefaultRetryIntervalMs int = 0
)

type HttpRequest struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
	context context.Context
}

func NewHTTPRequest() *HttpRequest {
	return &HttpRequest{Headers: make(map[string]string), context: context.Background()}
}

func (r *HttpRequest) SetMethod(method string) *HttpRequest {
	r.Method = method
	return r
}

func (r *HttpRequest) SetURL(url string) *HttpRequest {
	r.URL = url
	return r
}

func (r *HttpRequest) SetByteBody(body []byte) *HttpRequest {
	r.Body = body
	return r
}

func (r *HttpRequest) SetStringBody(body string) *HttpRequest {
	r.Body = []byte(body)
	return r
}

func (r *HttpRequest) SetHeader(header, value string) *HttpRequest {
	r.Headers[header] = value
	return r
}

func (r *HttpRequest) AddContext(ctx context.Context) *HttpRequest {
	r.context = ctx
	return r
}

func (r *HttpRequest) Build() (req *http.Request, err error) {
	if !(r.Method == http.MethodGet || r.Method == http.MethodPost) {
		return nil, errors.New("not found method")
	}

	var body io.Reader

	if r.Body != nil {
		body = bytes.NewBuffer(r.Body)
	}

	reqContext := context.Background()
	if r.context != nil {
		reqContext = r.context
	}

	req, err = http.NewRequestWithContext(reqContext, r.Method, r.URL, body)
	if err != nil {
		return nil, err
	}

	for header, value := range r.Headers {
		req.Header.Set(header, value)
	}

	return req, nil
}

type HttpResponse struct {
	Status int
	Header http.Header
	Body   io.ReadCloser
}

type HttpTransport struct {
	client        *http.Client
	maxRetryTimes int
	retryInterval time.Duration
}

func NewHTTPTransport(retryCount int, retryIntervalMs int) (*HttpTransport, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	return &HttpTransport{
		client:        client,
		maxRetryTimes: retryCount,
		retryInterval: time.Duration(retryIntervalMs) * time.Millisecond,
	}, nil
}

func NewHTTPTransportWithProxy(retryCount int, retryIntervalMs int, proxyUrl string) (*HttpTransport, error) {
	proxy, err := url.Parse(proxyUrl)
	if err != nil {
		return nil, errors.New("fail parse proxy url")
	}

	tr := &http.Transport{
		Proxy: http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{Transport: tr}

	return &HttpTransport{
		client:        client,
		maxRetryTimes: retryCount,
		retryInterval: time.Duration(retryIntervalMs) * time.Millisecond,
	}, nil
}

func (tr *HttpTransport) send(req *http.Request) (*HttpResponse, error) {
	resp, err := tr.client.Do(req)
	if err != nil {
		return nil, err
	}

	return &HttpResponse{
		Status: resp.StatusCode,
		Header: resp.Header,
		Body:   resp.Body,
	}, nil
}

func (tr *HttpTransport) Send(ctx context.Context, request *HttpRequest) (result *HttpResponse, err error) {
	if ctx == nil {
		return nil, errors.New("provided nil context")
	}

	req, err := request.AddContext(ctx).Build()
	if err != nil {
		return nil, err
	}

	for retryTimes := 0; retryTimes < tr.maxRetryTimes; retryTimes++ {
		result, err = tr.send(req)

		if err == nil {
			// check response code and on 500 range errors from server
			if !tr.isRetryStatusCode(result.Status) {
				break
			}

			// clear result body so we can reuse existing connection for next retry
			if err = tr.drainBody(result.Body); err != nil {
				break
			}
		}

		// check status of context. if context done - stop executing and return result
		if tr.isContextDone(ctx) {
			break
		}

		// wait some time to allow server to recover
		time.Sleep(tr.retryInterval)
	}

	return result, err
}

func (tr *HttpTransport) drainBody(body io.ReadCloser) error {
	defer body.Close()
	_, err := io.Copy(ioutil.Discard, io.LimitReader(body, respReadLimit))
	return err
}

func (tr *HttpTransport) isRetryStatusCode(status int) bool {
	return status == 0 || status >= 500
}

func (tr *HttpTransport) isContextDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
