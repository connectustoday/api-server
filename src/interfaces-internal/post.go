package interfaces_internal

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"log"
)

type IPost struct {
	SchemaVersion uint32        `bson:"schema_version"`
	ID            bson.ObjectId `bson:"_id"`
	Account       string        `bson:"account"`
	Content       string        `bson:"content"`
	CreatedAt     int64         `bson:"created_at"`
	ReplyTo       string        `bson:"reply_to"`
	Multimedia    IAttachment   `bson:"multimedia"`
	Tags          []string      `bson:"tags"`
	LikesCount    uint64        `bson:"likes_count"`
	CommentsCount uint64        `bson:"comments_count"`
	SharesCount   uint64        `bson:"shares_count"`
	Likes         []string      `bson:"likes"`
	Comments      []string      `bson:"comments"`
	Shares        []string      `bson:"shares"`
	Visibility    []string      `bson:"visibility"`
}

func InitIPostIndexes(collection *mgo.Collection) {
	err := collection.EnsureIndex(mgo.Index{
		Key: []string{"tags", "_id"},
		Background: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}