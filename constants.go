package hms

const (
	// VisibilityUnspecified for unspecified visibility
	VisibilityUnspecified = "VISIBILITY_UNSPECIFIED"
	// private visibility
	VisibilityPrivate = "PRIVATE"
	// public visibility
	VisibilityPublic = "PUBLIC"
	// secret visibility
	VisibilitySecret = "SECRET"
)

const (
	// high priority
	DeliveryPriorityHigh = "HIGH"
	// normal priority
	DeliveryPriorityNormal = "NORMAL"
)

const (
	// high priority
	NotificationPriorityHigh = "HIGH"
	// default priority
	NotificationPriorityDefault = "NORMAL"
	// low priority
	NotificationPriorityLow = "LOW"
)

const (
	// very low urgency
	UrgencyVeryLow = "very-low"
	// low urgency
	UrgencyLow = "low"
	// normal urgency
	UrgencyNormal = "normal"
	// high urgency
	UrgencyHigh = "high"
)

const (
	// webPush text direction auto
	DirAuto = "auto"
	// webPush text direction ltr
	DirLtr = "ltr"
	// webPush text direction rtl
	DirRtl = "rtl"
)

const (
	StyleBigText = iota + 1
)

const (
	TypeIntentOrAction = iota + 1
	TypeUrl
	TypeApp
	TypeRichResource
)

const (
	FastAppTargetDevelop = iota + 1
	FastAppTargetProduct
)

const (
	// test user
	TargetUserTypeTest = iota + 1
	// formal user
	TargetUserTypeFormal
	// VoIP user
	TargetUserTypeVoIP
)

const (
	// auth url
	authUrl = "https://login.cloud.huawei.com/oauth2/v2/token"

	// push server url
	sendMessageURLFmt = "https://api.push.hicloud.com/v1/%s/messages:send"
)
