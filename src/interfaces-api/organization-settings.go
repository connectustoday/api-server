package interfaces_api

type IOrganizationSettingsAPI struct {
	AllowMessagesFromUnknown bool   `json:"allow_messages_from_unknown"`
	EmailNotifications       bool   `json:"email_notifications"`
	Type                     string `json:"type"`
	// organization specific settings
	IsNonprofit bool `json:"is_nonprofit"`
}
