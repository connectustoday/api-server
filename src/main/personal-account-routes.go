package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/julienschmidt/httprouter"
	"interfaces-conv"
	"net/http"
)

// Get personal account profile route
// GET /v1/profile
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountProfileRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	FetchProfileRouteHelper(account, w)
}

// Get personal account route
// GET /v1/account
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountRoute(w http.ResponseWriter, _ *http.Request, _ httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	FetchAccountRouteHelper(account, w)
}

// Get personal account settings route
// GET /v1/settings
// https://connectustoday.github.io/api-server/api-reference#accounts

func GetPersonalAccountSettingsRoute(w http.ResponseWriter, _ *http.Request, p httprouter.Params, account bson.M) {
	w.Header().Set("Content-Type", "application/json")
	acc, err := interfaces_conv.ConvertBSONToIAccount(account)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	var b []byte
	if acc.Type == "User" {
		d, _ := interfaces_conv.ConvertBSONToIUser(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIUserSettingsAPI(d.Settings, d.Type))
	} else if acc.Type == "Organization" {
		d, _ := interfaces_conv.ConvertBSONToIOrganization(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIOrganizationSettingsAPI(d.Settings, d.Type))
	}
	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	_, err = w.Write([]byte(b))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}