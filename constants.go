package hms

import "time"

const (
	// auth url
	authUrl = "https://oauth-login.cloud.huawei.com/oauth2/v3/token"

	// push server url
	sendMessageURLFmt = "https://api.push.hicloud.com/v1/%s/messages:send"

	MaxMessageTTLSec = 15 * 24 * 60 * 60 // 15 days in seconds

	// default time for refreshing token
	RefreshTokenTime = time.Duration(30 * time.Minute)
)
