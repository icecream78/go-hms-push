package hms

const (
	// success code from push server
	SuccessCode = "80000000"

	// Some tokens are successfully sent. Tokens identified by illegal_token are those failed to be sent.
	SomeTokenSuccessErrorCode = "80100000"

	// parameter invalid code from push server
	ParameterErrorCode = "80100001"

	// The number of tokens must be 1 when a synchronization message is sent.
	SingleTokenSyncErrorCode = "80100002"

	// Incorrect message structure.
	IncorrectMessageErrorCode = "80100003"

	// The message expiration time is earlier   than the current time.
	ExpireTimeErrorCode = "80100004"

	// The collapse_key message field is invalid.
	CollapseKeyErrorCode = "80100013"

	// The message contains sensitive information.
	MessageInsecureErrorCode = "80100016"

	// OAuth authentication error.
	TokenFailedErrorCode = "80200001"

	// OAuth token expired.
	TokenTimeoutErrorCode = "80200003"

	// The current app does not have the permission to send push messages.
	NoPushPermissionErrorCode = "80300002"

	// All tokens are invalid.
	AllTokenInvalidErrorCode = "80300007"

	// The message body size exceeds the default value.
	BodyToBigErrorCode = "80300008"

	// The number of tokens in the message body exceeds the default value.
	TokensToMuchErrorCode = "80300010"

	// You are not authorized to send high-priority notification messages.
	NotAuthForHighPriorityMsgErrorCode = "80300011"

	// System internal error.
	InternalErrorCode = "81000001"
)

type HuaweiResponse struct {
	// Result code.
	Code string `json:"code"`

	// Result code description.
	Msg string `json:"msg"`

	// Request ID.
	RequestId string `json:"requestId"`
}
