package hms

import "errors"

type AndroidConfig struct {
	// Mode for the HUAWEI Push Kit server to control messages cached in user offline status.
	// These cached messages will be delivered once the user device goes online again.
	// The options are as follows:
	// 0: Only the latest offline message sent by each app to the user device is cached.
	// -1: All offline messages are cached.
	// 1-100: Offline message cache group ID. Offline messages are cached by group.
	//Each group can cache only one offline message for each app.
	// For example, if you send 10 messages and set collapse_key to 1 for the
	// first five messages and to 2 for the rest, the latest offline message
	// whose value of collapse_key is 1 and the latest offline message whose value of collapse_key is 2
	// are sent to the user.
	CollapseKey int `json:"collapse_key,omitempty"`

	// Delivery priority of a data message. The options can be HIGH and NORMAL
	// You need to apply for the permission when setting the parameter to HIGH,
	// in which the app process can be forcibly started when a data message reaches a user's mobile phone.
	// Please refer to https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-faq-v4#apply_special_permissions
	Urgency string `json:"urgency,omitempty"`

	// Scenario where a high-priority data message is sent.
	// Currently, this parameter can only be set to PLAY_VOICE (voice playing) and additional permission is required.
	// Please refer to https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-faq-v4#apply_special_permissions
	Category string `json:"category,omitempty"`

	// Message cache time, in seconds. When a user device is offline, the HUAWEI Push Kit server caches messages.
	// If the user device goes online within the message cache time, the messages are delivered.
	// Otherwise, the messages are discarded.
	// The default value is 86400 (1 day), and the maximum value is 1296000 (15 days).
	TTL string `json:"ttl,omitempty"`

	// Tag of a message in a batch delivery task.
	// The tag is returned to the app server when HUAWEI Push Kit sends the message receipt.
	// The app server can analyze message delivery statistics based on bi_tag.
	BiTag string `json:"bi_tag,omitempty"`

	// State of a mini program when a quick app sends a data message. The options are as follows:
	// 1: development state.
	// 2: production state (default value).
	FastAppTarget int `json:"fast_app_target,omitempty"`

	// Custom message payload. If the data parameter is set, the value of the message.data field is overwritten.
	Data string `json:"data,omitempty"`

	// Android notification message structure.
	Notification *AndroidNotification `json:"notification,omitempty"`
}

// AndroidNotification represents android notification params
type AndroidNotification struct {
	// Android notification message title. If the title parameter is set,
	// the value of the message.notification.title field is overwritten.
	// Before a message is sent, you must set at least one of title and message.notification.title.
	Title string `json:"title,omitempty"`

	// Android notification message body. If the body parameter is set,
	// the value of the message.notification.body field is overwritten.
	// Before a message is sent, you must set at least one of body and message.notification.body.
	Body string `json:"body,omitempty"`

	// Customized small icon on the left of a notification message.
	// The icon file must be stored in the /res/raw directory of an app.
	// For example, the value /raw/ic_launcher indicates the local icon file ic_launcher.xxx stored in /res/raw.
	// Currently, supported file formats include PNG and JPG.
	// For details about the specifications of the custom small icon, please refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-icon-spec
	Icon string `json:"icon,omitempty"`

	// Custom notification bar button colors in the #RRGGBB format,
	// where RR indicates the red hexadecimal color, GG indicates the green hexadecimal color,
	// and BB indicates the blue hexadecimal color.
	// Example: #FFEEFF
	Color string `json:"color,omitempty"`

	// Customized message notification ringtone, which is valid during channel creation.
	// The ringtone file must be stored in the /res/raw directory of an app.
	// For example, the value /raw/shake indicates the local ringtone file /res/raw/shake.xxx stored in /res/raw.
	// Currently, various file formats such as MP3, WAV, and MPEG are supported.
	// If this parameter is not set, the default system ringtone will be used.
	Sound string `json:"sound,omitempty"`

	// Indicates whether to use the default ringtone. The options are as follows:
	// true: The default ringtone is used.
	// false: A custom ringtone is used.
	DefaultSound bool `json:"default_sound,omitempty"`

	// Message tag. Messages that use the same message tag in the same app will be overwritten by the latest message.
	Tag string `json:"tag,omitempty"`

	// 	Message tapping action.
	// This parameter is mandatory for Android notification messages.
	ClickAction *ClickAction `json:"click_action,omitempty"`

	// ID in a string format of the localized message body
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	BodyLocKey string `json:"body_loc_key,omitempty"`

	// Variable parameter of the localized message body.
	// Example: ["args1","args2"]
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	BodyLocArgs []string `json:"body_loc_args,omitempty"`

	// ID in a string format of the localized message title
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	TitleLocKey string `json:"title_loc_key,omitempty"`

	// Variable parameter of the localized message title.
	// Example: ["args1","args2"]
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	TitleLocArgs []string `json:"title_loc_args,omitempty"`

	// Message in multiple languages. body_loc_key and title_loc_key are read from multi_lang_key first.
	// If they are not read from multi_lang_key, they will be read from the local character string of the APK.
	// A maximum of three languages can be set.
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	MultiLangKey map[string]interface{} `json:"multi_lang_key,omitempty"`

	// Customized channel for displaying notification messages.
	// Customized channels are supported in the Android O version or later.
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927575-1
	ChannelId string `json:"channel_id,omitempty"`

	// Brief description of an Android notification message.
	NotifySummary string `json:"notify_summary,omitempty"`

	// The URL of custom large icon on the right of an Android notification message.
	// The function is the same as that of the message.notification.image field.
	// If the image parameter is set, the value of the message.notification.image field is overwritten.
	// The URL must be an HTTPS URL, example: https://example.com/image.png
	Image string `json:"image,omitempty"`

	// Notification bar style. The options are as follows:
	// 0: default style.
	// 1: bigText.
	// 3: inbox style.
	Style int `json:"style,omitempty"`

	// Android notification message title in bigText style.
	// This parameter is mandatory when style is set to 1.
	// When the notification bar is displayed after big_title is set, big_title instead of title is used.
	BigTitle string `json:"big_title,omitempty"`

	// Android notification message body in bigText style.
	// This parameter is mandatory when style is set to 1.
	// When the notification bar is displayed after big_body is set, big_body instead of body is used.
	BigBody string `json:"big_body,omitempty"`

	// Message display duration, in milliseconds. Messages are automatically deleted after the duration expires.
	AutoClear int `json:"auto_clear,omitempty"`

	// Unique notification ID of a message.
	// If a message does not contain the ID or the ID is -1, NC will generate a unique ID for the message.
	// Different notification messages can use the same notification ID, so that new messages can overwrite old messages.
	NotifyId int `json:"notify_id,omitempty"`

	// Message group.
	// For example, if 10 messages are sent and the group parameter of the messages is set to 10,
	// only one message is displayed in the notification bar of the mobile phone.
	Group string `json:"group,omitempty"`

	// Android notification message badge control.
	Badge *BadgeNotification `json:"badge,omitempty,omitempty"`

	// Content displayed on the status bar after the device receives a notification message.
	Ticker string `json:"ticker,omitempty"`

	// Indicates whether an Android notification message is not still displayed in the notification bar
	// after a user taps the message. The options are as follows:
	// true: Yes.
	// false: No.
	AutoCancel bool `json:"auto_cancel,omitempty"`

	// Message sorting event. Android notification messages are sorted based on this value.
	// This event is displayed in the notification bar.
	// For example: 2014-10-02T15:01:23.045123456Z
	When string `json:"when,omitempty"`

	// Android notification message priority, which determines the message notification behavior of a user device.
	// The options are as follows:
	// LOW: common (silent) message
	// NORMAL: important message
	// HIGH: very important message
	Importance string `json:"importance,omitempty"`

	// Indicates whether to use the default vibration mode.
	UseDefaultVibrate bool `json:"use_default_vibrate,omitempty"`

	// Indicates whether to use the default breath light mode.
	UseDefaultLight bool `json:"use_default_light,omitempty"`

	// Custom vibration mode for an Android notification message.
	// Each array element adopts the format of [0-9]+|[0-9]+[sS]|[0-9]+[.][0-9]{1,9}|[0-9]+[.][0-9]{1,9}[sS].
	// For example, ["3.5S","2S","1S","1.5S"].
	// A maximum of ten array elements are supported.
	// The value of each element is an integer ranging from 0 to 60.
	VibrateConfig []string `json:"vibrate_config,omitempty"`

	// Android notification message visibility. The options are as follows:
	// VISIBILITY_UNSPECIFIED
	// PRIVATE
	// PUBLIC
	// SECRET
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#visibility
	Visibility string `json:"visibility,omitempty"`

	// Custom breath light mode.
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-References/push-sendapi#lightsettings
	LightSettings *LightSettings `json:"light_settings,omitempty"`

	// Indicates whether to display notification messages on the foreground when an app is running on the foreground
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-other#h1-1576146927576-2
	ForegroundShow bool `json:"foreground_show,omitempty"`

	// TODO: add inbox_content
	// TODO: add buttons
}

type ClickAction struct {
	// Message tapping action type. The options are as follows:
	// 1: custom tapping action.
	// 2: tap to open a specified URL.
	// 3: tap to start the app.
	// 4: tap to access rich media information.
	Type int `json:"type"`

	// Indicates which specific app`s page to open
	// When type is set to 1, you must set at least one of intent and action.
	// For details refer to: https://developer.huawei.com/consumer/en/doc/development/HMS-Guides/push-basic-capability#specified_page
	Intent string `json:"intent,omitempty"`

	// URL to be opened. The value must be an HTTPS URL, example: https://example.com/image.png.
	// This parameter is mandatory when type is set to 2.
	Action string `json:"action,omitempty"`

	// URL of a rich media resource to be opened. The value must be an HTTPS URL.
	// Rich media is a self-contained HTML resource package in .zip format.
	// This parameter is mandatory when type is set to 4.
	Url string `json:"url,omitempty"`

	// Action corresponding to the activity of the page to be opened when the custom app page is opened through the action.
	// When type is set to 1, you must set at least one of intent and action.
	RichResource string `json:"rich_resource,omitempty"`
}

type BadgeNotification struct {
	// Accumulative badge number, which is an integer ranging from 1 to 99.
	AddNum int `json:"add_num,omitempty"`

	// Class name in APK name + App entry activity format.
	// For example: com.huawei.codelabpush.MainActivity
	SetNum int `json:"set_num,omitempty"`

	// Set badge number, which is an integer ranging from 0 to 99.
	// If add_num and set_num exist at the same time, set_num is valid.
	Class string `json:"class,omitempty"`
}

type LightSettings struct {
	// Breath light color. This parameter is mandatory when light_settings is set.
	Color *Color `json:"color"`

	// Interval when a breath light is on, in the format of \d+|\d+[sS]|\d+.\d{1,9}|\d+.\d{1,9}[sS].
	// This parameter is mandatory when light_settings is set.
	LightOnDuration string `json:"light_on_duration,omitempty"`

	// Interval when a breath light is off, in the format of \d+|\d+[sS]|\d+.\d{1,9}|\d+.\d{1,9}[sS].
	// This parameter is mandatory when light_settings is set.
	LightOffDuration string `json:"light_off_duration,omitempty"`
}

type Color struct {
	// Alpha setting of the RGB color. The default value is 1, and the value range is [0,1].
	Alpha float32 `json:"alpha"`

	// Red setting of the RGB color. The default value is 0, and the value range is [0,1].
	Red float32 `json:"red"`

	// Green setting of the RGB color. The default value is 0, and the value range is [0,1].
	Green float32 `json:"green"`

	// Blue setting of the RGB color. The default value is 0, and the value range is [0,1].
	Blue float32 `json:"blue"`
}

func validateAndroidConfig(androidConfig *AndroidConfig) error {
	if androidConfig == nil {
		return nil
	}

	if androidConfig.CollapseKey < -1 || androidConfig.CollapseKey > 100 {
		return errors.New("collapse_key must be in interval [-1 - 100]")
	}

	if androidConfig.Urgency != "" &&
		androidConfig.Urgency != DeliveryPriorityHigh &&
		androidConfig.Urgency != DeliveryPriorityNormal {
		return errors.New("delivery_priority must be 'HIGH' or 'NORMAL'")
	}

	if androidConfig.TTL != "" && !ttlPattern.MatchString(androidConfig.TTL) {
		return errors.New("malformed ttl")
	}

	if androidConfig.FastAppTarget != 0 &&
		androidConfig.FastAppTarget != FastAppTargetDevelop &&
		androidConfig.FastAppTarget != FastAppTargetProduct {
		return errors.New("invalid fast_app_target")
	}

	// validate android notification
	return validateAndroidNotification(androidConfig.Notification)
}

func validateAndroidNotification(notification *AndroidNotification) error {
	if notification == nil {
		return nil
	}

	if notification.Sound == "" && notification.DefaultSound == false {
		return errors.New("sound must not be empty when default_sound is false")
	}

	if err := validateAndroidNotifyStyle(notification); err != nil {
		return err
	}

	if err := validateAndroidNotifyPriority(notification); err != nil {
		return err
	}

	if err := validateVibrateTimings(notification); err != nil {
		return err
	}

	if err := validateVisibility(notification); err != nil {
		return err
	}

	if err := validateLightSetting(notification); err != nil {
		return err
	}

	if notification.Color != "" && !colorPattern.MatchString(notification.Color) {
		return errors.New("color must be in the form #RRGGBB")
	}

	// validate click action
	return validateClickAction(notification.ClickAction)
}

func validateAndroidNotifyStyle(notification *AndroidNotification) error {
	switch notification.Style {
	case StyleBigText:
		if notification.BigTitle == "" {
			return errors.New("big_title must not be empty when style is 1")
		}

		if notification.BigBody == "" {
			return errors.New("big_body must not be empty when style is 1")
		}
	}
	return nil
}

func validateAndroidNotifyPriority(notification *AndroidNotification) error {
	if notification.Importance != "" &&
		notification.Importance != NotificationPriorityHigh &&
		notification.Importance != NotificationPriorityDefault &&
		notification.Importance != NotificationPriorityLow {
		return errors.New("Importance must be 'HIGH', 'NORMAL' or 'LOW'")
	}
	return nil
}

func validateVibrateTimings(notification *AndroidNotification) error {
	if notification.VibrateConfig != nil {
		if len(notification.VibrateConfig) > 10 {
			return errors.New("vibrate_timings can't be more than 10 elements")
		}
		for _, vibrateTiming := range notification.VibrateConfig {
			if !ttlPattern.MatchString(vibrateTiming) {
				return errors.New("malformed vibrate_timings")
			}
		}
	}
	return nil
}

func validateVisibility(notification *AndroidNotification) error {
	if notification.Visibility == "" {
		notification.Visibility = VisibilityPrivate
		return nil
	}
	if notification.Visibility != VisibilityUnspecified && notification.Visibility != VisibilityPrivate &&
		notification.Visibility != VisibilityPublic && notification.Visibility != VisibilitySecret {
		return errors.New("visibility must be VISIBILITY_UNSPECIFIED, PRIVATE, PUBLIC or SECRET")
	}
	return nil
}

func validateLightSetting(notification *AndroidNotification) error {
	if notification.LightSettings == nil {
		return nil
	}

	if notification.LightSettings.Color == nil {
		return errors.New("light_settings.color can't be nil")
	}

	if notification.LightSettings.LightOnDuration == "" ||
		!ttlPattern.MatchString(notification.LightSettings.LightOnDuration) {
		return errors.New("light_settings.light_on_duration is empty or malformed")
	}

	if notification.LightSettings.LightOffDuration == "" ||
		!ttlPattern.MatchString(notification.LightSettings.LightOffDuration) {
		return errors.New("light_settings.light_off_duration is empty or malformed")
	}
	return nil
}

func validateClickAction(clickAction *ClickAction) error {
	if clickAction == nil {
		return errors.New("click_action object must not be null")
	}

	switch clickAction.Type {
	case TypeIntentOrAction:
		if clickAction.Intent == "" && clickAction.Action == "" {
			return errors.New("at least one of intent and action is not empty when type is 1")
		}
	case TypeUrl:
		if clickAction.Url == "" {
			return errors.New("url must not be empty when type is 2")
		}
	case TypeApp:
	case TypeRichResource:
		if clickAction.RichResource == "" {
			return errors.New("rich_resource must not be empty when type is 4")
		}
	default:
		return errors.New("type must be in the interval [1 - 4]")
	}
	return nil
}

func GetDefaultAndroid() *AndroidConfig {
	android := &AndroidConfig{
		Urgency:      DeliveryPriorityNormal,
		TTL:          "86400s",
		Notification: nil,
	}
	return android
}

func GetDefaultAndroidNotification() *AndroidNotification {
	notification := &AndroidNotification{
		DefaultSound: true,
		Importance:   NotificationPriorityDefault,
		ClickAction:  getDefaultClickAction(),
	}

	notification.UseDefaultVibrate = true
	notification.UseDefaultLight = true
	notification.Visibility = VisibilityPrivate
	notification.ForegroundShow = true

	notification.AutoCancel = true

	return notification
}

func getDefaultClickAction() *ClickAction {
	return &ClickAction{
		Type:   TypeIntentOrAction,
		Action: "Action",
	}
}

func GetDefaultAndroidNotificationMessage(tokenArr []string) *HuaweiMessage {
	msg := NewNotificationMsgRequest()
	msg.Message.Token = tokenArr
	msg.Message.Android = GetDefaultAndroid()
	msg.Message.Android.Notification = GetDefaultAndroidNotification()
	msg.Message.Android.Notification.Body = "Notification body text"
	return msg
}
