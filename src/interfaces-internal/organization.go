package interfaces_internal

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type IValidations struct {
	UserID       string `bson:"user_id"`
	ExperienceID string `bson:"experience_id"`
}

type IOrganization struct {
	SchemaVersion        int                   `bson:"schema_version"`
	ID                   bson.ObjectId         `bson:"_id" json:"id"`
	Email                string                `bson:"email"`
	Password             string                `bson:"password"`
	OAuthToken           string                `bson:"oauth_token"`
	OAuthService         string                `bson:"oauth_service"`
	IsEmailVerified      bool                  `bson:"is_email_verified"`
	LastLogin            int64                 `bson:"last_login"`
	Notifications        []INotification       `bson:"notifications"`
	Avatar               string                `bson:"avatar"`
	Header               string                `bson:"header"`
	CreatedAt            int64                 `bson:"created_at"`
	PendingConnections   []string              `bson:"pending_connections"`
	RequestedConnections []string              `bson:"requested_connections"`
	Posts                []string              `bson:"posts"`
	Liked                []ICom                `bson:"liked"`
	Shared               []ICom                `bson:"shared"`
	Settings             IOrganizationSettings `bson:"settings"`
	AdminNote            string                `bson:"admin_note"`
	Type                 string                `bson:"type"`
	PasswordResetToken   string                `bson:"password_reset_token" json:"password_reset_token"`
	VerifyEmailToken     string                `bson:"verify_email_token" json:"verify_email_token"`
	AuthKey              string                `bson:"auth_key" json:"auth_key"`
	// organization specific fields
	PreferredName         string               `bson:"preferred_name"`
	IsVerified            bool                 `bson:"is_verified"`
	Opportunities         []string             `bson:"opportunities"`
	OrgInfo               IOrganizationProfile `bson:"org_info"`
	ExperienceValidations []IValidations       `bson:"experience_validations"`
}

func InitIOrganizationIndexes(collection *mgo.Collection) {
	err := collection.EnsureIndex(mgo.Index{
		Key:        []string{"preferred_name"},
		Background: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}
