package hms

const (
	// auth url
	authUrl = "https://login.cloud.huawei.com/oauth2/v2/token"

	// push server url
	sendMessageURLFmt = "https://api.push.hicloud.com/v1/%s/messages:send"

	MaxMessageTTLSec = 15 * 24 * 60 * 60 // 15 days in seconds
)
