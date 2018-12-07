package interfaces_internal

type IUser struct {
	SchemaVersion        int             `bson:"schema_version"`
	UserName             string          `bson:"username"`
	Email                string          `bson:"email"`
	Password             string          `bson:"password"`
	OAuthToken           string          `bson:"oauth_token"`
	OAuthService         string          `bson:"oauth_service"`
	IsEmailVerified      bool            `bson:"is_email_verified"`
	LastLogin            int64           `bson:"last_login"`
	Notifications        []INotification `bson:"notifications"`
	Avatar               string          `bson:"avatar"`
	Header               string          `bson:"header"`
	CreatedAt            int64           `bson:"created_at"`
	PendingConnections   []string        `bson:"pending_connections"`
	RequestedConnections []string        `bson:"requested_connections"`
	Posts                []string        `bson:"posts"`
	Liked                []ICom          `bson:"liked"`
	Shared               []ICom          `bson:"shared"`
	Settings             interface{}     `bson:"settings"`
	AdminNote            string          `bson:"admin_note"`
	Type                 string          `bson:"type"`
	// IUser specific fields
	FirstName    string       `bson:"first_name"`
	MiddleName   string       `bson:"middle_name"`
	LastName     string       `bson:"last_name"`
	Birthday     string       `bson:"birthday"`
	Gender       string       `bson:"gender"`
	PersonalInfo IUserProfile `bson:"personal_info"`
	Experiences []IExperience `bson:"experiences"`
}
