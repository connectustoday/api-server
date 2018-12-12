package interfaces_api

type IUserSettingsAPI struct {
	AllowMessagesFromUnknown bool   `json:"allow_messages_from_unknown"`
	EmailNotifications       bool   `json:"email_notifications"`
	Type                     string `json:"type"`
	// user specific settings
	IsFullNameVisible bool     `json:"is_full_name_visible"`
	BlockedUsers      []string `json:"blocked_users"`
}
