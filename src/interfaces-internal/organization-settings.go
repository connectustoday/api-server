package interfaces_internal

type IOrganizationSettings struct {
	AllowMessagesFromUnknown bool   `bson:"allow_messages_from_unknown"`
	EmailNotifications       bool   `bson:"email_notifications"`
	// organization specific settings
	IsNonprofit bool `bson:"is_nonprofit"`
}
