package interfaces_internal

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type ICom struct {
	Post string `bson:"post"`
	When int    `bson:"when"`
}

type IAccount struct {
	SchemaVersion        int              `bson:"schema_version" json:"schema_version"`
	ID                   bson.ObjectId    `bson:"_id" json:"id"`
	UserName             string           `bson:"username" json:"username"`
	Email                string           `bson:"email" json:"email"`
	Password             string           `bson:"password" json:"password"`
	OAuthToken           string           `bson:"oauth_token" json:"oauth_token"`
	OAuthService         string           `bson:"oauth_service" json:"oauth_service"`
	IsEmailVerified      bool             `bson:"is_email_verified" json:"is_email_verified"`
	LastLogin            int64            `bson:"last_login" json:"last_login"`
	Notifications        []INotification  `bson:"notifications" json:"notifications"`
	Avatar               string           `bson:"avatar" json:"avatar"`
	Header               string           `bson:"header" json:"header"`
	CreatedAt            int64            `bson:"created_at" json:"created_at"`
	PendingConnections   []string         `bson:"pending_connections" json:"pending_connections"`
	RequestedConnections []string         `bson:"requested_connections" json:"requested_connections"`
	Posts                []string         `bson:"posts" json:"posts"`
	Liked                []ICom           `bson:"liked" json:"liked"`
	Shared               []ICom           `bson:"shared" json:"shared"`
	Settings             IAccountSettings `bson:"settings" json:"settings"`
	AdminNote            string           `bson:"admin_note" json:"admin_note"`
	Type                 string           `bson:"type" json:"type"`
	PasswordResetToken   string           `bson:"password_reset_token" json:"password_reset_token"`
	VerifyEmailToken     string           `bson:"verify_email_token" json:"verify_email_token"`
	AuthKey              string           `bson:"auth_key" json:"auth_key"`
}

func InitIAccountIndexes(collection *mgo.Collection) {
	err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"username", "email", "password", "avatar", "admin_note"},
		Background: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}
