package hms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Transporter interface {
	Send(ctx context.Context, req *HttpRequest) (*HttpResponse, error)
}

type HuaweiClient struct {
	appId     string
	appSecret string

	tokenInfo *TokenMsg

	isAutoUpdateRun bool

	transport Transporter

	mu sync.RWMutex
}

// Token return current token value
func (c *HuaweiClient) Token() *TokenMsg {
	c.mu.RLock()
	token := c.tokenInfo
	c.mu.RUnlock()

	return token
}

// StartAutoUpdateToken runs logic for auto regeneration token
func (c *HuaweiClient) StartAutoUpdateToken(ctx context.Context, updateTime time.Duration) {
	c.isAutoUpdateRun = true
	forceRefreshToken := true

	for {
		if forceRefreshToken {
			if _, err := c.RequestToken(ctx); err != nil {
				time.Sleep(10 * time.Second) // myabe make it more progressive
				continue
			}

			forceRefreshToken = false
			continue
		}

		select {
		case <-time.After(updateTime):
			forceRefreshToken = true
		case <-ctx.Done():
			c.isAutoUpdateRun = false
			return
		}
	}
}

// RequestToken sends manual request to huawei cloud and updates token after successful response
func (c *HuaweiClient) RequestToken(ctx context.Context) (*TokenMsg, error) {
	token, err := c.requestToken(ctx)
	if err != nil {
		return nil, ErrorRefreshToken
	}

	c.mu.Lock()
	c.tokenInfo = token
	c.mu.Unlock()

	return token, nil
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

	var token *TokenMsg

	if !c.isAutoUpdateRun {
		reqCtx, reqCtxCancel := context.WithTimeout(ctx, time.Second*3)
		defer reqCtxCancel()

		tokenMsg, err := c.RequestToken(reqCtx)
		if err != nil {
			return nil, err
		}

		token = tokenMsg
	} else {
		token = c.Token()
	}

	request := NewHTTPRequest().
		SetMethod(http.MethodPost).
		SetURL(fmt.Sprintf(sendMessageURLFmt, c.appId)).
		SetByteBody(body).
		SetHeader("Content-Type", "application/json;charset=utf-8").
		SetHeader("Authorization", "Bearer "+token.AccessToken)

	resp, err := c.sendHttpRequest(ctx, request)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func (c *HuaweiClient) requestToken(ctx context.Context) (*TokenMsg, error) {
	u, _ := url.Parse(c.appSecret)
	body := fmt.Sprintf("grant_type=client_credentials&client_secret=%s&client_id=%s", u.String(), c.appId)

	request := NewHTTPRequest().
		SetMethod(http.MethodPost).
		SetURL(authUrl).
		SetStringBody(body).
		SetHeader("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.transport.Send(ctx, request)
	if err != nil {
		return nil, err
	}

	if resp.Status != http.StatusOK {
		return nil, errors.New("fail get token")
	}

	respDecoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var token TokenMsg
	if err := respDecoder.Decode(&token); err != nil {
		return nil, err
	}

	return &token, nil
}

func (c *HuaweiClient) sendHttpRequest(ctx context.Context, request *HttpRequest) (*HuaweiResponse, error) {
	resp, err := c.transport.Send(ctx, request)
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
