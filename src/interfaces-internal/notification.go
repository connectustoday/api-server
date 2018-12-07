package interfaces_internal

import "github.com/globalsign/mgo/bson"

type INotification struct {
	ID        bson.ObjectId `bson:"_id"`
	CreatedAt int64         `bson:"created_at"`
	Type      string        `bson:"type"`
	Content   string        `bson:"content"`
	Account   string        `bson:"account"`
}
