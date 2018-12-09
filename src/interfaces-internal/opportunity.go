package interfaces_internal

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type IOpportunity struct {
	SchemaVersion        uint32        `bson:"schema_version"`
	ID                   bson.ObjectId `bson:"_id"`
	Organization         string        `bson:"organization"`
	Name                 string        `bson:"name"`
	Description          string        `bson:"description"`
	Address              IAddress      `bson:"address"`
	IsSignupsEnabled     bool          `bson:"is_signups_enabled"`
	NumberOfPeopleNeeded int64         `bson:"number_of_people_needed"`
	Tags                 []string      `bson:"tags"`
	InterestedUsers      []string      `bson:"interested_users"`
	ShiftTimes           []string      `bson:"shift_times"`
	MethodOfContact      []string      `bson:"method_of_contact"`
	CreatedAt            int64         `bson:"created_at"`
}

func InitIOpportunityIndexes(collection *mgo.Collection) {
	err := collection.EnsureIndex(mgo.Index{
		Key: []string{"organization", "name", "is_signups_enabled", "_id"},
		Background: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}