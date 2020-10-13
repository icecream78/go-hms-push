package hms

import (
	"errors"
	"time"
)

type WebPushConfig struct {
	// WebPush message header
	Headers *WebPushHeaders `json:"headers,omitempty"`

	// WebPush notification message structure
	Notification *WebPushNotification `json:"notification,omitempty"`

	// WebPush agent parameter
	HmsOptions *HmsWebPushOption `json:"hms_options,omitempty"`
}

type WebPushHeaders struct {
	// Message cache time, in seconds, for example, 20, 20s, or 20S.
	TTL string `json:"ttl,omitempty"`

	// Message ID, which can be used to overwrite undelivered messages.
	Topic string `json:"topics,omitempty"`

	// Message emergency level. The value can only be very-low, low, normal, or high.
	Urgency string `json:"urgency,omitempty"`
}

type HmsWebPushOption struct {
	// Default URL for redirection when no action is performed.
	Link string `json:"link,omitempty"`
}

type WebPushNotification struct {
	// Title of a web app notification message.
	// If the title parameter is set, the value of the message.notification.title field is overwritten.
	// Before a message is sent, you must set at least one of title and message.notification.title.
	Title string `json:"title,omitempty"`

	// Body of a web app notification message.
	// If the body parameter is set, the value of the message.notification.body field is overwritten.
	// Before a message is sent, you must set at least one of body and message.notification.body.
	Body string `json:"body,omitempty"`

	// Small icon URL.
	Icon string `json:"icon,omitempty"`
	// Large image URL.
	Image string `json:"image,omitempty"`

	// Language.
	Lang string `json:"lang,omitempty"`

	// Notification message group tag. Multiple same tags are collapsed and the latest one is displayed.
	// This function is used only for mobile phone browsers.
	Tag string `json:"tag,omitempty"`

	// Browser icon URL, which only applies to mobile phone browsers and is used to replace the default browser icon.
	Badge string `json:"badge,omitempty"`

	// Text direction, which can be set to auto, ltr, or rtl.
	Dir string `json:"dir,omitempty"`

	// Vibration interval, in milliseconds. The value is an integer by default. The value range is [100,200,300].
	Vibrate []int `json:"vibrate,omitempty"`

	// Message reminding flag.
	Renotify bool `json:"renotify,omitempty"`

	// Indicates that notification messages should remain active until a user taps or closes them.
	RequireInteraction bool `json:"require_interaction,omitempty"`

	// Message sound-free and vibration-free reminding flag.
	Silent bool `json:"silent,omitempty"`

	// Sending timestamp.
	Timestamp int64 `json:"timestamp,omitempty"`

	// Message action.
	Actions []*WebPushAction `json:"actions,omitempty"`
}

type WebPushAction struct {
	// Action name.
	Action string `json:"action,omitempty"`

	// URL for the button icon of an action.
	Icon string `json:"icon,omitempty"`

	// Title of an action.
	Title string `json:"title,omitempty"`
}

func GetDefaultWebPushConfig() *WebPushConfig {
	return &WebPushConfig{
		Headers:    getDefaultHeaders(),
		HmsOptions: getDefaultHmsOptions(),
	}
}

func getDefaultHeaders() *WebPushHeaders {
	return &WebPushHeaders{
		TTL:     "990",
		Topic:   "topic",
		Urgency: UrgencyVeryLow,
	}
}

func getDefaultHmsOptions() *HmsWebPushOption {
	return &HmsWebPushOption{
		Link: "https://www.huawei.com",
	}
}

func GetDefaultWebNotification() *WebPushNotification {
	return &WebPushNotification{
		Dir:       DirAuto,
		Silent:    true,
		Timestamp: time.Now().Unix(),
	}
}

func validateWebPushConfig(webPushConfig *WebPushConfig) error {
	if webPushConfig == nil {
		return nil
	}

	if err := validateWebPushHeaders(webPushConfig.Headers); err != nil {
		return err
	}

	return validateWebPushNotification(webPushConfig.Notification)
}

func validateWebPushHeaders(headers *WebPushHeaders) error {
	if headers == nil {
		return nil
	}

	if headers.Urgency != "" &&
		headers.Urgency != UrgencyHigh &&
		headers.Urgency != UrgencyNormal &&
		headers.Urgency != UrgencyLow &&
		headers.Urgency != UrgencyVeryLow {
		return errors.New("priority must be 'high', 'normal', 'low' or 'very-low'")
	}
	return nil
}

func validateWebPushNotification(notification *WebPushNotification) error {
	if notification == nil {
		return nil
	}

	if err := validateWebPushAction(notification.Actions); err != nil {
		return err
	}

	if err := validateWebPushDirection(notification.Dir); err != nil {
		return err
	}
	return nil
}

func validateWebPushAction(actions []*WebPushAction) error {
	if actions == nil {
		return nil
	}

	for _, action := range actions {
		if action.Action == "" {
			return errors.New("web common action can't be empty")
		}
	}
	return nil
}

func validateWebPushDirection(dir string) error {
	if dir != DirAuto && dir != DirLtr && dir != DirRtl {
		return errors.New("web common dir must be 'auto', 'ltr', 'rtl'")
	}
	return nil
}

func GetDefaultWebNotificationMessage(tokenArr []string) *HuaweiMessage {
	msg := NewNotificationMsgRequest()
	msg.Message.Token = tokenArr
	msg.Message.WebPush = GetDefaultWebPushConfig()
	return msg
}
