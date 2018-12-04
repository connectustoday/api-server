package api_server

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-api"
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
	query := IAccountCollection.Find(bson.M{"username": p.ByName("id")})
	count, err := query.Count()
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError + " (Problem finding account)", 3002)
	}
	if count == 0 {
		SendError(w, http.StatusNotFound, notFound + " (Account not found)", 3003)
	}
	account := interfaces_internal.IAccount{}
	err = query.One(&account)
	accountapi := interfaces_api.ConvertToIAccountAPI(account)

	w.Write([]byte())
}

func GetAccountProfileRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func GetAccountConnectionsRoute(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

