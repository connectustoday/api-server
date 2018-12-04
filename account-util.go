package api_server

import (
	"github.com/globalsign/mgo/bson"
	"interfaces-internal"
	"log"
)

func VerifyUniqueUsername(username string) bool {
	var results []interfaces_internal.IAccount
	err := IAccountCollection.Find(bson.M{"username": username}).All(&results)
	if err != nil {
		log.Print(err)
	}
	return len(results) == 0
}
