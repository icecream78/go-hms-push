package hms

import "errors"

var (
	// client errors
	ErrorFailGetHttpClient = errors.New("failed to get http client")
	ErrorAppIdEmpty        = errors.New("appId can't be empty")
	ErrorEmptyTransport    = errors.New("passed empty transport")
	ErrorRefreshToken      = errors.New("refresh token fail")

	// transport errors
	ErrorUnsupportedMethod = errors.New("not found method")
	ErrorFailParseProxyURL = errors.New("fail parse proxy url")

	// android errors
	ErrorCollapseKeyValue      = errors.New("collapse_key must be in interval [-1 - 100]")
	ErrorSoundEmpty            = errors.New("sound must not be empty when default_sound is false")
	ErrorColorFormat           = errors.New("color must be in the form #RRGGBB")
	ErrorBigTitleEmpty         = errors.New("big_title must not be empty when style is 1")
	ErrorBigBodyEmpty          = errors.New("big_body must not be empty when style is 1")
	ErrorVibrateTimingOverflow = errors.New("vibrate_timings can't be more than 10 elements")
	ErrorVibrateTimingDuration = errors.New("vibrate_timings are more 60 seconds")
	ErrorLightThemeColorNil    = errors.New("light_settings.color can't be nil")
	ErrorClickActionNil        = errors.New("click_action object must not be null")
	ErrorIntentAndActionEmpty  = errors.New("at least one of intent and action is not empty when type is 1")
	ErrorEmptyURL              = errors.New("url must not be empty when type is 2")
	ErrorRichResourceEmpty     = errors.New("rich_resource must not be empty when type is 4")
	ErrorClickActionValue      = errors.New("click_action type must be in the interval [1 - 4]")

	// web errors
	ErrorWebActionEmpty = errors.New("web common action can't be empty")

	// common message
	ErrorInvalidVisabilityType           = errors.New("invalid visibility type")
	ErrorInvalidDeliveryPriorityType     = errors.New("invalid delivery priority type")
	ErrorInvalidNotificationPriorityType = errors.New("invalid notification priority type")
	ErrorInvalidUrgencyType              = errors.New("invalid urgency type")
	ErrorInvalidTextDirectionType        = errors.New("invalid text direction type")
	ErrorInvalidNotificationBarStyleType = errors.New("invalid notification bar style type")
	ErrorInvalidClickActionType          = errors.New("invalid click action type")
	ErrorInvalidFastAppStateType         = errors.New("invalid fast app state type")
)
