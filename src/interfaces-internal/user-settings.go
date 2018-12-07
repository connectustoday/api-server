package interfaces_internal

type IUserSettings struct {
	AllowMessagesFromUnknown bool   `bson:"allow_messages_from_unknown"`
	EmailNotifications       bool   `bson:"email_notifications"`
	// user specific settings
	IsFullNameVisible bool     `bson:"is_full_name_visible"`
	BlockedUsers      []string `bson:"blocked_users"`
}
