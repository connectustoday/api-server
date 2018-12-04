package interfaces_internal

type IAccountSettings struct {
	AllowMessagesFromUnknown bool `bson:"allow_messages_from_unknown"`
	EmailNotifications       bool `bson:"email_notifications"`
}
