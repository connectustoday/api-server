package interfaces_internal

import "github.com/globalsign/mgo/bson"

type When struct {
	Begin string `bson:"begin"`
	End   string `bson:"end"`
}

type IExperience struct {
	ID            bson.ObjectId `bson:"_id"`
	SchemaVersion uint32        `bson:"schema_version"`
	Location      IAddress      `bson:"address"`
	Name          string        `bson:"name"`
	Organization  string        `bson:"organization"`
	Opportunity   string        `bson:"opportunity"`
	Description   string        `bson:"description"`
	When          When          `bson:"when"`
	IsVerified    bool          `bson:"is_verified"`
	EmailVerify   bool          `bson:"email_verify"`
	CreatedAt     int64         `bson:"created_at"`
	Hours         int64         `bson:"hours"`
	EmailJWT      string        `bson:"emailjwt"`
}
