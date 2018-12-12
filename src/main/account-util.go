package main

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"interfaces-conv"
	"net/http"
)

func GetAccountDBAsBSON(query bson.M) (bson.M, error) {
	var account bson.M
	err := IAccountCollection.Find(query).One(&account)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func FetchAccountRouteHelper(account bson.M, w http.ResponseWriter) {
	acc, err := interfaces_conv.ConvertBSONToIAccount(account)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	var b []byte
	if acc.Type == "User" {
		d, _ := interfaces_conv.ConvertBSONToIUser(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIUserAPI(d))
	} else if acc.Type == "Organization" {
		d, _ := interfaces_conv.ConvertBSONToIOrganization(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIOrganizationAPI(d))
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

func FetchProfileRouteHelper(account bson.M, w http.ResponseWriter) {
	acc, err := interfaces_conv.ConvertBSONToIAccount(account)

	if err != nil {
		SendError(w, http.StatusInternalServerError, internalServerError+" (Problem finding account)", 4001)
		return
	}

	var b []byte
	if acc.Type == "User" {
		d, _ := interfaces_conv.ConvertBSONToIUser(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIUserProfileAPI(d.PersonalInfo, d.Type))
	} else if acc.Type == "Organization" {
		d, _ := interfaces_conv.ConvertBSONToIOrganization(account)
		b, err = json.Marshal(interfaces_conv.ConvertToIOrganizationProfileAPI(d.OrgInfo, d.Type))
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