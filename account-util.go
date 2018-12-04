package api_server

import (
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-internal"
	"log"
	"net/http"
)

func VerifyUniqueUsername(username string) bool {
	var results []interfaces_internal.IAccount
	err := IAccountCollection.Find(bson.M{"username": username}).All(&results)
	if err != nil {
		log.Print(err)
	}
	return len(results) == 0
}

func RegisterRoute(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func VerifyEmailRequestRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetAccountRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}