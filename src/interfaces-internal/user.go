package interfaces_internal

import (
	"github.com/globalsign/mgo"
	"log"
)

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
	Settings             IUserSettings   `bson:"settings"`
	AdminNote            string          `bson:"admin_note"`
	Type                 string          `bson:"type"`
	PasswordResetToken   string          `bson:"password_reset_token" json:"password_reset_token"`
	VerifyEmailToken     string          `bson:"verify_email_token" json:"verify_email_token"`
	// IUser specific fields
	FirstName    string        `bson:"first_name"`
	MiddleName   string        `bson:"middle_name"`
	LastName     string        `bson:"last_name"`
	Birthday     string        `bson:"birthday"`
	Gender       string        `bson:"gender"`
	PersonalInfo IUserProfile  `bson:"personal_info"`
	Experiences  []IExperience `bson:"experiences"`
}

func InitIUserIndexes(collection *mgo.Collection) {
	err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"first_name"},
		Background: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}
