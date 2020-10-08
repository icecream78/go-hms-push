package hms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type Transporter interface {
	Send(ctx context.Context, req *HttpRequest) (*HttpResponse, error)
}

type TokenMsg struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Scope            string `json:"scope"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type HuaweiClient struct {
	appId     string
	token     string
	appSecret string
	client    Transporter
}

// NewClient creates a instance of the huawei cloud common client
// It's contained in huawei cloud app and provides service through huawei cloud app
func NewHuaweiClient(appId, appSecret string) (*HuaweiClient, error) {
	if appId == "" {
		return nil, errors.New("appId can't be empty")
	}

	client, err := NewHTTPTransport(DefaultRetryCount, DefaultRetryIntervalMs)
	if err != nil {
		return nil, errors.New("failed to get http client")
	}

	return &HuaweiClient{
		appId:     appId,
		appSecret: appSecret,
		client:    client,
	}, nil
}

func NewHuaweiClientWithTransport(appId, appSecret string, transport Transporter) (*HuaweiClient, error) {
	client, err := NewHuaweiClient(appId, appSecret)
	if err != nil {
		return nil, err
	}

	if err := client.SetTransport(transport); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *HuaweiClient) SetTransport(transport Transporter) error {
	if transport == nil {
		return errors.New("passed empty transport")
	}
	c.client = transport

	return nil
}

// GetToken return current token value
func (c *HuaweiClient) GetToken() string {
	return c.token
}

func (c *HuaweiClient) requestToken(ctx context.Context) (string, error) {
	u, _ := url.Parse(c.appSecret)
	body := fmt.Sprintf("grant_type=client_credentials&client_secret=%s&client_id=%s", u.String(), c.appId)

	request := NewHTTPRequest().
		SetMethod(http.MethodPost).
		SetURL(authUrl).
		SetStringBody(body).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Send(ctx, request)
	if err != nil {
		return "", err
	}

	if resp.Status != http.StatusOK {
		return "", errors.New("fail get token")
	}

	respDecoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var token TokenMsg
	if err := respDecoder.Decode(&token); err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (c *HuaweiClient) refreshToken(ctx context.Context) error {
	token, err := c.requestToken(ctx)
	if err != nil {
		return errors.New("refresh token fail")
	}

	c.token = token
	return nil
}

func (c *HuaweiClient) executeApiOperation(ctx context.Context, request *HttpRequest) (*HuaweiResponse, error) {
	resp, err := c.sendHttpRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	// if need to retry for token timeout or other reasons
	retry, err := c.isNeedRetry(ctx, resp)
	if err != nil {
		return nil, err
	}

	if retry {
		return c.sendHttpRequest(ctx, request)
	}
	return resp, err
}

func (c *HuaweiClient) sendHttpRequest(ctx context.Context, request *HttpRequest) (*HuaweiResponse, error) {
	resp, err := c.client.Send(ctx, request)
	if err != nil {
		return nil, err
	}

	respDecoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var hr HuaweiResponse
	if err := respDecoder.Decode(&hr); err != nil {
		return nil, err
	}

	return &hr, nil
}

// if token is timeout or error or other reason, need to refresh token and send again
func (c *HuaweiClient) isNeedRetry(ctx context.Context, resp *HuaweiResponse) (bool, error) {
	if !(resp.Code == TokenTimeoutErrorCode || resp.Code == TokenFailedErrorCode) {
		return false, nil
	}

	if err := c.refreshToken(ctx); err != nil {
		return false, err
	}

	return true, nil
}

// SendMessage sends a message to huawei cloud common
// One of Token, Topic and Condition fields must be invoked in message
// If validationOnly is set to true, the message can be verified by not sent to users
func (c *HuaweiClient) SendMessage(ctx context.Context, msgRequest *HuaweiMessage) (*HuaweiResponse, error) {
	if err := msgRequest.Validate(); err != nil {
		return nil, err
	}

	body, err := json.Marshal(msgRequest)
	if err != nil {
		return nil, err
	}

	// initial send call after client init
	if c.token == "" {
		if err := c.refreshToken(ctx); err != nil {
			return nil, err
		}
	}

	request := NewHTTPRequest().
		SetMethod(http.MethodPost).
		SetURL(fmt.Sprintf(sendMessageURLFmt, c.appId)).
		SetByteBody(body).
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Authorization", "Bearer "+c.token)

	resp, err := c.executeApiOperation(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, err
}
