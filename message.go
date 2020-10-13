package hms

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

var (
	colorPattern = regexp.MustCompile("^#[0-9a-fA-F]{6}$")
)

type TTL struct {
	t time.Duration
}

func NewTTL(dur time.Duration) *TTL {
	return &TTL{
		t: dur,
	}
}

func (t TTL) Seconds() float64 {
	return t.t.Seconds()
}

func (t TTL) MarshalJSON() ([]byte, error) {
	sec := t.Seconds()
	if sec > MaxMessageTTLSec {
		sec = MaxMessageTTLSec
	}

	s := fmt.Sprintf("%gS", sec)
	return []byte(s), nil
}

// HuaweiMessage represents list of request params and payload for push api
type HuaweiMessage struct {
	// ValidateOnly indicates whether a message is test or not.
	// Test message is only used to verify format validity and
	// not pushed to user device
	ValidateOnly bool `json:"validate_only"`

	// Push message structure
	Message *Message `json:"message"`
}

// TODO: think how rewrite this validation in proper way
func (hr *HuaweiMessage) Validate() error {
	// validate field target, one of Token, Topic and Condition must be invoked
	if err := validateFieldTarget(hr.Message.Token, hr.Message.Topic, hr.Message.Condition); err != nil {
		return err
	}

	// validate android config
	if err := validateAndroidConfig(hr.Message.Android); err != nil {
		return err
	}

	// validate web common config
	if err := validateWebPushConfig(hr.Message.WebPush); err != nil {
		return err
	}
	return nil
}

// TODO: understand how this function works on what purpose of that stuff
func validateFieldTarget(token []string, strings ...string) error {
	count := 0
	if token != nil {
		count++
	}

	for _, s := range strings {
		if s != "" {
			count++
		}
	}

	if count == 1 {
		return nil
	}
	return errors.New("token, topics or condition must be choice one")
}

type Message struct {
	// Custom message payload, which can be a common string or a string in JSON format.
	// Example: "your data" or "{'param1':'value1','param2':'value2'}"
	// If the message body contains message.data and does not contain message.notification or
	// message.android.notification, the message is a data message.
	Data string `json:"data,omitempty"`

	// Notification message content
	Notification *Notification `json:"notification,omitempty"`

	// Android push message control
	// This parameter is mandatory for Android notification messages.
	Android *AndroidConfig `json:"android,omitempty"`

	// iOS push message control
	// This parameter is mandatory for iOS messages.
	// block is empty, because corresponding documentation is missing
	Apns interface{} `json:"apns,omitempty"`

	// Web app push message control.
	// This parameter is mandatory for iOS messages.
	WebPush *WebPushConfig `json:"webpush,omitempty"`

	// Push token of the target user of a message. You must set one of token, topic, and condition.
	// Example:   ["pushtoken1","pushtoken2"]
	Token []string `json:"token,omitempty"`

	// Topic subscribed by the target user of a message.
	// (Currently, this parameter only applies to Android apps).
	// You must set one of token, topic, and condition.
	Topic string `json:"topic,omitempty"`

	// Condition (topic combination expression) for sending a message to the target user.
	// (Currently, this parameter applies only to Android apps).
	// You must set one of token, topic, and condition.
	//
	// A Boolean expression of target topics can be specified to send messages based on a combination of condition expressions.
	// Supported boolean operations:
	// '&&': logical AND
	// '||': logical OR
	// '! ': logical negative
	// '()': priority control
	// 'in': keywords
	// A maximum of five topics can be included in a condition expression.
	// "'TopicA' in topics && ('TopicB' in topics || 'TopicC' in topics)"
	// The preceding expression indicates that messages are sent to devices that subscribe to topics A and B or topic C. Devices that subscribe to a single topic do not receive the messages.
	Condition string `json:"condition,omitempty"`
}

type Notification struct {
	// 	Notification message title.
	Title string `json:"title,omitempty"`

	// Notification message content.
	Body string `json:"body,omitempty"`

	// The URL of custom large icon on the right of a notification message.
	// If this parameter is not set, the icon is not displayed.
	// The URL must be an HTTPS URL, example: https://example.com/image.png
	Image string `json:"image,omitempty"`
}

// NewNotificationMsgRequest will return a new MessageRequest instance with default value to send notification message.
// developers should set at least on of Message.Token or  Message.Topic or Message.Condition
func NewNotificationMsgRequest() *HuaweiMessage {
	return &HuaweiMessage{
		ValidateOnly: false,
		Message: &Message{
			Data: "This is a transparent message data",
			Notification: &Notification{
				Title: "notification title",
				Body:  "This is a notification message body",
			},
		},
	}
}
