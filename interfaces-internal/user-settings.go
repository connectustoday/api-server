package interfaces_internal

type IUserSettings struct {
	*IAccountSettings
	IsFullNameVisible bool     `bson:"is_full_name_visible"`
	BlockedUsers      []string `bson:"blocked_users"`
}
