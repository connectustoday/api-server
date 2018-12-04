package interfaces_internal

type ICom struct {
	Post string `bson:"post"`
	When int    `bson:"when"`
}

type IAccount struct {
	SchemaVersion        int              `bson:"schema_version"`
	UserName             string           `bson:"username"`
	Email                string           `bson:"email"`
	Password             string           `bson:"password"`
	OAuthToken           string           `bson:"oauth_token"`
	OAuthService         string           `bson:"oauth_service"`
	IsEmailVerified      bool             `bson:"is_email_verified"`
	LastLogin            int64            `bson:"last_login"`
	Notifications        []INotification  `bson:"notifications"`
	Avatar               string           `bson:"avatar"`
	Header               string           `bson:"header"`
	CreatedAt            int64            `bson:"created_at"`
	PendingConnections   []string         `bson:"pending_connections"`
	RequestedConnections []string         `bson:"requested_connections"`
	Posts                []string         `bson:"posts"`
	Liked                []ICom           `bson:"liked"`
	Shared               []ICom           `bson:"shared"`
	Settings             IAccountSettings `bson:"settings"`
	AdminNote            string           `bson:"admin_note"`
}
