package main

import (
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Get personal account settings route
// GET /v1/profile
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountProfileRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	FetchProfileRouteHelper(account, w)
}